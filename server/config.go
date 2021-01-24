/*
 * @Author: Vongola
 * @FilePath: /JFFun/server/config.go
 * @Date: 2021-01-23 23:35:12
 * @Description: file content
 * @描述: 文件描述
 * @LastEditTime: 2021-01-24 21:59:17
 * @LastEditors: Vongola
 */

package server

import (
	"flag"
	"strings"
)

func parseFlag() (release bool, cfgs []string) {
	r := flag.Bool(`r`, true, `发布模式运行`)
	c := flag.String(`c`, ``, `服务器配置文件，多服务器模式配置文件以";"号隔开`)
	flag.Parse()
	return *r, strings.Split(*c, ";")
}

type serverConfig struct {
	Name    string            `toml:"name" json:"name" yaml:"name"`
	Modules map[string]string `toml:"module" json:"module" yaml:"module"`
}

type moduleConfig struct {
	AutoRestart bool `toml:"autoRestart" json:"autoRestart" yaml:"autoRestart"`
	Buffer      int  `toml:"buffer" json:"buffer" yaml:"buffer"`
}
