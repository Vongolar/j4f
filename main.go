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
	"fmt"
	"j4f/core/module"
	mscheduler "j4f/core/modules/scheduler"
	"j4f/core/server"
	"j4f/modules/mlog"
)

func main() {
	fmt.Println(`Just For Fun`)

	l := new(mlog.M_Log)
	run(l)

	fmt.Println(`BYE`)
}

func run(mods ...module.Module) {
	server.Run(new(mscheduler.M_Scheduler), mods...)
}

//go:generate go generate ./proto/
