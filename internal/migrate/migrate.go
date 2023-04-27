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

func (m Migrate) Configs() []*GitConfigMap {
	return m.configs
}

type GitConfigMap struct {
	iniFile        *ini.File
	CurrentOrigin  Origin
	ReplacedOrigin Origin
	File           string
}

type Origin string

func (o Origin) String() string {
	return string(o)
}

func (o Origin) Replace(find, replace string) string {
	return strings.Replace(string(o), find, replace, 1)
}

func (m Migrate) ReplaceConfigOrigin(parentDir string) error {
	return filepath.Walk(parentDir, m.WalkFunc)
}

func (m Migrate) WalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !info.IsDir() && strings.Contains(path, GIT_CONFIG_FILE) {
		// fmt.Println(path, info.Size())
		m.log.Debugf("found git config path: %s\n", path)
		cfg, err := ini.Load(path)
		if err != nil {
			return err
		}
		gm := GitConfigMap{iniFile: cfg, File: path}

		if err := gm.SetNewConfig(path, m.find, m.replace); err != nil {
			m.log.Debugf("falied to set new value in git config path: %s\n", path)
			return err
		}

		m.configs = append(m.configs, &gm)

		return nil
	}
	return nil
}

func (g GitConfigMap) SetNewConfig(path, find, replace string) error {
	url := Origin(g.iniFile.Section(INI_SECTION).Key(INI_PROPERTY).Value())

	g.iniFile.Section(INI_SECTION).Key(INI_PROPERTY).SetValue(url.Replace(find, replace))
	return g.iniFile.SaveTo(g.File)
}

// func (g GitConfigMap) SetRemoteUrlOrigin(find, replace string) {
// 	originUrl := Origin(g.iniFile.Section(INI_SECTION).Key(INI_PROPERTY).Value())
// 	g.CurrentOrigin = originUrl
// 	newUrl := Origin(originUrl)
// 	newUrl.Replace(find, replace)
// 	g.ReplacedOrigin = newUrl
// }
