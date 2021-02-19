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
	"j4f/core/module"
	"j4f/core/server"
)

func main() {

	server1 := map[string]module.Module{}
	server2 := map[string]module.Module{}

	server.RunServers([]map[string]module.Module{server1, server2})

	// server.RunServer(server1)
}

//go:generate go generate ./cmd/
