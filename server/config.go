package jserver

import (
	jconfig "JFFun/serialization/config"
	"flag"
	"os"
)

type config struct {
	Modules map[string]moduleConfig
	Mysql   map[string]mysqlConfig
	Redis   map[string]redisConfig
}

type moduleConfig struct {
	Buf   int
	Paths []string
}

type mysqlConfig struct {
	User     string
	Password string
	Addr     string
}

type redisConfig struct {
	Password string
	DB       int
	Addr     string
}

var cfg config

var configPath string
var serverConfigFile = `server.toml`

func parseFlag() error {
	flag.StringVar(&configPath, "cfg", "./configs/", "server config file")
	flag.Parse()
	if _, err := os.Stat(configPath + serverConfigFile); err != nil {
		return err
	}
	return nil
}

func loadCfg() error {
	err := parseFlag()
	if err != nil {
		return err
	}

	return jconfig.LoadConfig(configPath+serverConfigFile, &cfg)
}
