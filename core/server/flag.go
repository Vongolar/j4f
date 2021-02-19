/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-19 14:35:20
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\server\flag.go
 * @Date: 2021-02-19 14:18:57
 * @描述: 文件描述
 */

package server

import (
	"flag"
	"strings"
)

type startupParameter struct {
	release     bool
	configFiles []string
}

func parseFlag() startupParameter {
	r := flag.Bool(`r`, true, `release mode`)
	s := flag.String(`c`, ``, `server config file`)

	flag.Parse()

	pars := &startupParameter{
		release:     *r,
		configFiles: parseConfigFiles(*s),
	}
	return *pars
}

func parseConfigFiles(s string) []string {
	res := strings.Split(s, ";")
	for i := len(res) - 1; i >= 0; i-- {
		if len(strings.TrimSpace(res[i])) == 0 {
			res = append(res[:i], res[i+1:]...)
		}
	}
	return res
}
