/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-04 19:01:11
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
)

func main() {
	server.MutliRun([]module.Module{new(http.M_Http)}, []module.Module{new(http.M_Http)}, []module.Module{new(http.M_Http)})
}
