/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-22 16:24:08
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\server\config.go
 * @Date: 2021-01-15 10:14:10
 * @描述: 服务器配置
 */
package server

import (
	"flag"
	"strings"
)

func parseFlag() (bool, []string) {
	release := flag.Bool("r", true, `release 模式运行`)
	cfg := flag.String("c", "", `服务器配置文件，多个服务器用';'隔开`)

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
