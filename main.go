/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-22 17:55:41
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\main.go
 * @Date: 2021-01-15 10:14:10
 * @描述: 文件描述
 */
package main

import (
	"JFFun/gate"
	"JFFun/jlog"
	"JFFun/module"
	"JFFun/server"
)

func main() {
	jlog.Info("Just For Fun")
	server.RunServers(
		[]module.Module{new(gate.M_Gate), new(gate.M_Gate), new(gate.M_Gate)},
		[]module.Module{new(gate.M_Gate), new(gate.M_Gate)},
		[]module.Module{new(gate.M_Gate)},
	)
	// server.Run(new(gate.M_Gate), new(gate.M_Gate), new(gate.M_Gate))
	jlog.Info("Good Bye")
}

//go:generate go generate ./proto
