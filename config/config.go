package config

import (
	"JFFun/serialize/toml"
	"flag"
)

var configPath string

//ParseServerConfig 解析服务器启动配置
func ParseServerConfig() {
	flag.StringVar(&configPath, "p", "./configs/", "配置文件目录")
	flag.Parse()
}

//ParseModuleConfig 解析服务器各模块配置
func ParseModuleConfig(name string, config interface{}) error {
	return toml.DecodeFile(configPath+name+".toml", config)
}
