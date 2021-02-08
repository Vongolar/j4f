/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 17:10:27
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
	"j4f/modules/http"
	"j4f/modules/schedule"
)

func main() {
	server.MutliRun([]module.Module{new(schedule.M_Schedule), new(http.M_Http)}, []module.Module{new(schedule.M_Schedule), new(http.M_Http)}, []module.Module{new(schedule.M_Schedule), new(http.M_Http)})
}

//go:generate go generate ./gen/
