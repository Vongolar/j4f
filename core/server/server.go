/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-19 16:58:09
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\server\server.go
 * @Date: 2021-02-19 11:56:00
 * @描述: 文件描述
 */

package server

import (
	"fmt"
	jconfig "j4f/core/config"
	"j4f/core/module"
	"os"
	"os/signal"
	"sync"
)

const (
	exitCode_noServer                 = 1
	exitCode_serverConfigCountNoMatch = 2
	exitCode_serverPanic              = 3
)

func RunServer(mods map[string]module.Module) {
	RunServers([]map[string]module.Module{mods})
}

func RunServers(servers []map[string]module.Module) {
	fmt.Println(`hello`)

	startupPar = parseFlag()

	if len(startupPar.configFiles) != len(servers) {
		//TODO:配置数目不一致
		os.Exit(exitCode_serverConfigCountNoMatch)
	}

	if len(servers) == 0 {
		//TODO:没有配置
		os.Exit(exitCode_noServer)
	}

	if len(servers) == 1 {
		defer func() {
			if err := recover(); err != nil {
				//TODO:服务器意外退出
				fmt.Println(`server panic`, err)
				os.Exit(exitCode_serverPanic)
			}
		}()
		run(startupPar.configFiles[0], servers[0])
	} else {
		var wg sync.WaitGroup
		hasErr := false
		for i, cfg := range startupPar.configFiles {
			sc, mods := cfg, servers[i]
			wg.Add(1)
			go func() {
				defer func() {
					if err := recover(); err != nil {
						//TODO:服务器意外退出
						fmt.Println(`server panic`, err)
						hasErr = true
					}
					wg.Done()
				}()
				run(sc, mods)
			}()
		}
		wg.Wait()
		if hasErr {
			os.Exit(exitCode_serverPanic)
		}
	}

	fmt.Println(`bye`)
}

var startupPar startupParameter

type server struct {
	cfg config
}

func run(cfg string, mods map[string]module.Module) {
	s := new(server)
	err := jconfig.ParseFile(cfg, &s.cfg)
	if err != nil {
		panic(err)
	}

	cSignal := make(chan os.Signal)
	signal.Notify(cSignal, os.Interrupt, os.Kill)
	<-cSignal
}
