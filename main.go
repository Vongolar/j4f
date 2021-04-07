/*
 * @Author: Vongola
 * @LastEditTime: 2021-04-07 18:28:29
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \j4f\main.go
 * @Date: 2021-02-04 11:42:36
 * @描述: 文件描述
 */
package main

import (
	"j4f/core/module"
	"j4f/core/modules/console"
	mscheduler "j4f/core/modules/scheduler"
	"j4f/core/server"
	mhttp "j4f/modules/gate/http"
	"j4f/modules/gate/socks5"
	"j4f/modules/mlog"
)

func main() {
	server.Info(`Just For Fun`)

	l := new(mlog.M_Log)
	s5 := new(socks5.M_Socks5)
	h := new(mhttp.M_Http)
	sin := new(console.M_Console)

	run(sin, l, h, s5)

	server.Info(`BYE`)
}

func run(mods ...module.Module) {
	server.Run(new(mscheduler.M_Scheduler), mods...)
}

//go:generate go generate ./proto/
