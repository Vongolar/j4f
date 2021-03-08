package server

import "j4f/core/loglevel"

var defaultConfig = config{
	Release:     true,
	MinLogLevel: loglevel.INFO,
}

type config struct {
	Release     bool           `toml:"release" json:"release" yaml:"release"`
	MinLogLevel loglevel.Level `toml:"minLogLevel" json:"minLogLevel" yaml:"minLogLevel"`

	ModuleConfigs []moduleConfig `toml:"modules" json:"modules" yaml:"modules"`
}

type moduleConfig struct {
	Name   string `toml:"name" json:"name" yaml:"name"`
	Config string `toml:"config" json:"config" yaml:"config"`
}
