package server

import (
	Jtoml "JFFun/serialization/toml"
	"flag"
	"io/ioutil"
	"os"
)

var configFile string
var cfg config

func parseFlag() error {
	flag.StringVar(&configFile, "cfg", "./configs/", "server config file")
	flag.Parse()
	configFile += `server.toml`
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return err
	}
	return nil
}

type config struct {
	Gate []string
}

func parseConfig() error {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	return Jtoml.Unmarshal(b, &cfg)
}
