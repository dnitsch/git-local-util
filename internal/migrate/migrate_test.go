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
func helperCreateTmpFsObject(t *testing.T, input string, randDirs []string) (string, []string) {
	createdConfig := []string{}

	dir, _ := os.MkdirTemp("", "git-config")
	for _, v := range randDirs {
		// os.Chdir(dir)
		bf := filepath.Join(dir, v, ".git") // , "config")
		if err := os.MkdirAll(bf, 0755); err != nil {
			t.Fatal(err)
		}
		file := filepath.Join(bf, "config")
		if err := os.WriteFile(file, []byte(fmt.Sprintf(input, v)), 0777); err != nil {
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
				tpl := `[core]
				repositoryformatversion = 0
[remote "origin"]
	url   = git@ssh.foo.com:v3/org1/someproj/%s
	fetch = +refs/heads/*:refs/remotes/origin/*`
				parentDir, files := helperCreateTmpFsObject(t, tpl, input)
				return parentDir, files, func() {
					_ = os.Remove(parentDir)
				}
			},
			[]string{"foo", "bar"},
		},
		"no directories found": {
			func(input []string) (string, []string, func()) {
				parentDir, files := helperCreateTmpFsObject(t, ``, input)
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
			outLog := &bytes.Buffer{}
			g := gmo.New("org1", newRepo, log.New(outLog, log.DebugLvl))
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

func Test(t *testing.T) {
	ttests := map[string]struct {
		tpl string
	}{
		"incorrectly formatted input": {
			`[co
url   = git@ssh.foo.com:v3/org1/someproj/%s
fetch = +refs/heads/*:refs/remotes/origin/*`,
		},
		// 		"key not found in config": {
		// 			`[core]
		// 				repositoryformatversion = 0
		// [notfound]
		// 	wrong   = git@ssh.foo.com:v3/org1/someproj/%s
		// 	fetch = +refs/heads/*:refs/remotes/origin/*`,
		// 		},
	}
	for name, tt := range ttests {
		t.Run(name, func(t *testing.T) {
			parentDir, _ := helperCreateTmpFsObject(t, tt.tpl, []string{"rm"})
			defer func() {
				os.Remove(parentDir)
			}()

			outLog := &bytes.Buffer{}
			// os.Remove(filepath.Join(parentDir, "rm"))
			g := gmo.New("org1", newRepo, log.New(outLog, log.DebugLvl))
			err := g.ReplaceConfigOrigin(parentDir)
			if err == nil {
				t.Errorf("got: <nil>, expected err to be thrown")
			}
		})
	}
}
