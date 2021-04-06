/*
 * @Author: Vongola
 * @LastEditTime: 2021-03-30 16:58:21
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: /JFFun/main.go
 * @Date: 2021-02-04 11:42:36
 * @描述: 文件描述
 */
package main

import (
	"j4f/core/module"
	mscheduler "j4f/core/modules/scheduler"
	"j4f/core/server"
	mhttp "j4f/modules/gate/http"
)

func main() {
	server.Info(`Just For Fun`)

	// l := new(mlog.M_Log)
	h := new(mhttp.M_Http)

	run(h)

	server.Info(`BYE`)
}

func run(mods ...module.Module) {
	server.Run(new(mscheduler.M_Scheduler), mods...)
}

//go:generate go generate ./proto/
