package config

import (
	"flag"

	toml "github.com/BurntSushi/toml"
)

func DecodeConfigFile(file string, config interface{}) error {
	_, err := toml.DecodeFile(file, config)
	return err
}

var configPath string

func ParseServerConfig() {
	flag.StringVar(&configPath, "p", "./configs/", "配置文件目录")
	flag.Parse()
}

func ParseModuleConfig(name string, config interface{}) error {
	_, err := toml.DecodeFile(configPath+name+".toml", config)
	return err
}
