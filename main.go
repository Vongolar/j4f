/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-19 17:02:08
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\main.go
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

	run(new(mlog.M_Log))

	fmt.Println(`BYE`)

	// server1 := map[string]module.Module{}
	// server2 := map[string]module.Module{}

	// server.RunServers([]map[string]module.Module{server1, server2})

	// // server.RunServer(server1)
}

func run(mods ...module.Module) {
	server.Run(new(mscheduler.M_Scheduler), mods...)
}

//go:generate go generate ./proto/
