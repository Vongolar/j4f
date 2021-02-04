/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-04 19:36:53
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\server\config.go
 * @Date: 2021-02-04 14:18:43
 * @描述: 文件描述
 */
package server

import (
	"flag"
	"strings"
)

type Conifg struct {
	Name    string         `toml:"name" yaml:"name" json:"name"`
	Modules []ModuleConfig `toml:"module" yaml:"module" json:"module"`
}

type args struct {
	release bool
	configs []string
}

func parseFlag() *args {
	res := new(args)
	flag.BoolVar(&res.release, "r", true, `发布模式`)
	cfgs := flag.String("c", "", "服务器配置")
	flag.Parse()

	res.configs = parseConfigs(*cfgs)

	return res
}

func parseConfigs(str string) []string {
	var res []string
	cfgPaths := strings.Split(str, ";")
	for _, cfg := range cfgPaths {
		cfg = strings.TrimSpace(cfg)
		if len(cfg) == 0 {
			continue
		}
		res = append(res, cfg)
	}
	return res
}
