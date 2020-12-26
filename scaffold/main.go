package main

import (
	"flag"
	"strings"
)

/*
 * @Author: Vongola
 * @FilePath: /JFFun/scaffold/main.go
 * @Date: 2020-12-26 15:57:02
 * @Description: Scaffold for 'Just For Fun'
 * @描述: Just For Fun 框架脚手架
 * @LastEditTime: 2020-12-26 16:33:04
 * @LastEditors: Vongola
 */

func main() {
	flag.Parse()

	flag.Arg(0)
}

var flags = map[string]int{
	"help": 0,
	"gen":  1,
}

func matchCommand(cmd string) bool {
	if _, ok := flags[strings.ToLower(cmd)]; ok {
		return true
	}
	return false
}

const defaultDes = `just4fun is a scaffold of JFFun.

Usage:
	just4fun <command> [arguments]

The commands are:
	gen			auto gen files

Use "just4fun help <command>" for more info about the command.`
