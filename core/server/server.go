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
	jlog "j4f/core/log"
	"j4f/core/module"
	"log"
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
	jlog.SetBuffLog()

	log.Println(`Just For Fun`)

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
		run(startupPar.configFiles[0], servers[0])
	} else {
		var wg sync.WaitGroup
		hasErr := false
		for i, cfg := range startupPar.configFiles {
			sc, mods := cfg, servers[i]
			wg.Add(1)
			go func() {
				defer wg.Done()
				if re := run(sc, mods); re != nil {
					hasErr = true
				}
			}()
		}
		wg.Wait()
		if hasErr {
			os.Exit(exitCode_serverPanic)
		}
	}

	log.Println(`BYE`)
}

var startupPar startupParameter

type server struct {
	cfg config
}

func run(cfg string, mods map[string]module.Module) (runErr error) {
	s := new(server)
	runErr = jconfig.ParseFile(cfg, &s.cfg)
	if runErr != nil {
		log.Fatalln(fmt.Errorf("%s config file parse error.", cfg), runErr)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(fmt.Errorf("server \"%s\" panic.", s.cfg.Name), err)
			runErr = fmt.Errorf("server panic: %v", err)
		}
	}()

	cSignal := make(chan os.Signal)
	signal.Notify(cSignal, os.Interrupt, os.Kill)
	<-cSignal
	return
}
