/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-24 21:42:27
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: /JFFun/main.go
 * @Date: 2021-01-15 10:14:10
 * @描述: 文件描述
 */
package main

import (
	"JFFun/jlog"
	"JFFun/module"
	"JFFun/module/template"
	"JFFun/server"
)

func main() {
	jlog.Info("Just For Fun")

	s1 := map[string]module.Module{
		"t1": new(template.MTemplate),
	}
	s2 := map[string]module.Module{
		"t1": new(template.MTemplate),
		"t2": new(template.MTemplate),
	}
	s3 := map[string]module.Module{
		"t1": new(template.MTemplate),
		"t2": new(template.MTemplate),
		"t3": new(template.MTemplate),
	}

	server.RunServers(s1, s2, s3)

	jlog.Info("Good Bye")
}

//go:generate go generate ./proto
