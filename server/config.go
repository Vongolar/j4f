package server

import (
	Jtoml "JFFun/serialization/toml"
	"flag"
	"io/ioutil"
	"os"
)

var configPath string
var cfg config

func parseFlag() error {
	flag.StringVar(&configPath, "cfg", "./configs/", "server config file")
	flag.Parse()
	if _, err := os.Stat(configPath + `server.toml`); os.IsNotExist(err) {
		return err
	}
	return nil
}

type config struct {
	Gate []string
}

func parseConfig() error {
	b, err := ioutil.ReadFile(configPath + `server.toml`)
	if err != nil {
		return err
	}
	return Jtoml.Unmarshal(b, &cfg)
}
