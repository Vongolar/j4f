/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-04 18:50:14
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\gen\genConfig\main.go
 * @Date: 2021-02-04 16:44:56
 * @描述: 文件描述
 */

package main

import (
	"flag"
	"fmt"
	"j4f/core/config"
	"j4f/core/schedule"
	"j4f/core/server"
	"os"
	"path/filepath"
)

var serverConfig = server.Conifg{
	Name: "server1",
	Modules: []schedule.ModuleConfig{
		{Name: "HTTP1", Buffer: 5},
		{Name: "HTTP2", Buffer: 10},
		{Name: "HTTP3", Buffer: 15},
	},
}

func main() {
	genConfigFile(serverConfig)
}

func genConfigFile(cfg interface{}) {
	path := flag.String("out", "", `输出文件`)
	flag.Parse()
	ext := filepath.Ext(*path)

	os.Remove(*path)
	file, err := os.OpenFile(*path, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	config.Encode(ext, file, cfg)
}

//go:generate go run ./main.go -out ../../config/server.toml
