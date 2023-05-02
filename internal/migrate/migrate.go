package migrate

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/dnitsch/simplelog"
	"gopkg.in/ini.v1"
)

const (
	INI_SECTION     = `remote "origin"`
	INI_PROPERTY    = "url"
	GIT_CONFIG_FILE = ".git/config"
)

type Migrate struct {
	configs       []*GitConfigMap
	log           log.Loggeriface
	find, replace string
}

func New(find, replace string, log log.Loggeriface) Migrate {
	return Migrate{
		configs: []*GitConfigMap{},
		find:    find,
		replace: replace,
		log:     log,
	}
}

type GitConfigMap struct {
	iniFile        *ini.File
	CurrentOrigin  Origin
	ReplacedOrigin Origin
	File           string
}

type Origin string

func (o Origin) Replace(find, replace string) string {
	return strings.Replace(string(o), find, replace, 1)
}

// ReplaceConfigOrigin identifies all .git/configs in a directory recursively
// loads the INI config file and attemps to perform a replacement
func (m Migrate) ReplaceConfigOrigin(parentDir string) error {
	m.log.Debugf("attempting to read parentDir: %s", parentDir)
	return filepath.WalkDir(parentDir, m.walkFunc)
}

func (m Migrate) walkFunc(path string, info os.DirEntry, err error) error {
	if err != nil {
		m.log.Errorf("error reading dir in walkDirFunc: %v", err)
		return err
	}

	if !info.IsDir() && strings.Contains(path, GIT_CONFIG_FILE) {
		// fmt.Println(path, info.Size())
		m.log.Debugf("found git config path: %s\n", path)
		cfg, err := ini.Load(path)
		if err != nil {
			m.log.Errorf("error loading ini file from path: %v", err)
			return err
		}
		gm := GitConfigMap{iniFile: cfg, File: path}

		if err := gm.setNewConfig(path, m.find, m.replace); err != nil {
			m.log.Debugf("falied to set new value in git config path: %s\n", path)
			return err
		}
		m.log.Debugf("successfully replaced origin in: %s, exchanged '%s' for '%s'", path, m.find, m.replace)

		return nil
	}
	return nil
}

func (g GitConfigMap) setNewConfig(path, find, replace string) error {
	url := Origin(g.iniFile.Section(INI_SECTION).Key(INI_PROPERTY).Value())
	g.iniFile.Section(INI_SECTION).Key(INI_PROPERTY).SetValue(url.Replace(find, replace))
	return g.iniFile.SaveTo(g.File)
}
