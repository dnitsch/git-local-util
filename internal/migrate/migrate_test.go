package migrate_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	gmo "github.com/dnitsch/git-local-util/internal/migrate"
	log "github.com/dnitsch/simplelog"
)

var newRepo = "org-new"

// helperCreateTmpFsObject returns a the parent temp dir
// and a list of locations where fake configs were created
func helperCreateTmpFsObject(t *testing.T, randDirs []string) (string, []string) {
	createdConfig := []string{}
	tpl := `[core]
    repositoryformatversion = 0
    filemode                = true
    bare                    = false
    logallrefupdates        = true
    ignorecase              = true
    precomposeunicode       = true

[remote "origin"]
    url   = git@ssh.foo.com:v3/org1/someproj/%s
    fetch = +refs/heads/*:refs/remotes/origin/*`

	dir, _ := os.MkdirTemp("", "git-config")
	for _, v := range randDirs {
		// os.Chdir(dir)
		bf := filepath.Join(dir, v, ".git") // , "config")
		if err := os.MkdirAll(bf, 0755); err != nil {
			t.Fatal(err)
		}
		file := filepath.Join(bf, "config")
		if err := os.WriteFile(file, []byte(fmt.Sprintf(tpl, v)), 0777); err != nil {
			t.Fatal(err)
		}
		createdConfig = append(createdConfig, file)
	}
	return dir, createdConfig
}

func Test_ReplaceConfigOrigin_should_successfully_pass(t *testing.T) {
	ttests := map[string]struct {
		// parentDir func() string
		testDir func([]string) (string, []string, func())
		want    []string
	}{
		"when using absolute path": {
			func(input []string) (string, []string, func()) {
				parentDir, files := helperCreateTmpFsObject(t, input)
				return parentDir, files, func() {
					_ = os.Remove(parentDir)
				}
			},
			[]string{"foo", "bar"},
		},
		"no directories found": {
			func(input []string) (string, []string, func()) {
				parentDir, files := helperCreateTmpFsObject(t, input)
				return parentDir, files, func() {
					_ = os.Remove(parentDir)
				}
			},
			[]string{},
		},
	}
	for name, tt := range ttests {
		t.Run(name, func(t *testing.T) {
			parentDir, setupFiles, cleanUp := tt.testDir(tt.want)
			defer cleanUp()
			logOut := &bytes.Buffer{}
			g := gmo.New("org1", newRepo, log.New(logOut, log.DebugLvl))
			err := g.ReplaceConfigOrigin(parentDir)
			if err != nil {
				t.Fatal(err)
			}
			for _, file := range setupFiles {
				got, _ := os.ReadFile(file)
				if !strings.Contains(string(got), newRepo) {
					t.Errorf("got: %v\nexpected output to include: %v", string(got), newRepo)
				}
			}
		})
	}
}
