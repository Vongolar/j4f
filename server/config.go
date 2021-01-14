package server

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func parseFlag() (bool, []string) {
	release := flag.Bool("r", true, `release 模式`)
	cfg := flag.String("cfg", "", `服务器配置文件，多个服务器用';'隔开`)

	flag.Parse()
	return *release, strings.Split(*cfg, ";")
}

type config struct {
	Name    string             `json:"name" toml:"name" yaml:"name"`
	Modules []moduleConfigFile `json:"modules" toml:"modules" yaml:"modules"`
}

type moduleConfigFile struct {
	Name string `json:"name" toml:"name" yaml:"name"`
	Path string `json:"path" toml:"path" yaml:"path"`
}

func (cfg *config) checkModuleConfigExist(root string) error {
	for i, mc := range cfg.Modules {
		if filepath.IsAbs(mc.Path) {
			f, err := os.OpenFile(mc.Path, os.O_RDONLY, os.ModePerm)
			if err != nil {
				return fmt.Errorf("open module %s's config '%s' error : %v", mc.Name, mc.Path, err)
			}
			f.Close()
			continue
		}

		f, err := os.OpenFile(mc.Path, os.O_RDONLY, os.ModePerm)
		if err == nil {
			f.Close()
			continue
		}

		f, err = os.OpenFile(filepath.Join(root, mc.Path), os.O_RDONLY, os.ModePerm)
		if err != nil {
			return fmt.Errorf("open module %s's config '%s' error : %v", mc.Name, mc.Path, err)
		}
		f.Close()
		cfg.Modules[i].Path = filepath.Join(root, mc.Path)
	}
	return nil
}
