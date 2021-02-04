/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-04 19:36:35
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\server\server.go
 * @Date: 2021-02-04 11:30:35
 * @描述: 文件描述
 */

package server

import (
	"context"
	"fmt"
	"j4f/core/config"
	"j4f/core/log"
	"j4f/core/module"
	"os"
	"os/signal"
	"sync"
)

func Run(modules ...module.Module) {
	MutliRun(modules)
}

func MutliRun(servers ...[]module.Module) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	var serverList []*server
	for _, mods := range servers {
		if len(mods) == 0 {
			continue
		}

		s := &server{
			schedule: newSchedule(ctx, &wg),
		}

		serverList = append(serverList, s)
	}

	if len(serverList) == 0 {
		cancel()
		return
	}

	args := parseFlag()

	if len(args.configs) != len(serverList) {
		log.InfoTag(`server`, `服务器配置数目和启动服务器数目不匹配`)
		cancel()
		return
	}

	if args.release && len(serverList) > 1 {
		log.WarnTag(`server`, `发布模式不建议启动多个server`)
	}

	for i, s := range serverList {
		if err := config.DecodeConfigFromFile(args.configs[i], &s.cfg); err != nil {
			log.ErrorTag(`server`, fmt.Sprintf("服务器配置文件 %s 解析失败", args.configs[i]), err)
			cancel()
			return
		}

		if len(s.cfg.Modules) != len(servers[i]) {
			log.InfoTag(`server`, `服务器模块配置数目和启动服务模块数目不匹配`)
			cancel()
			return
		}

		for j, mc := range s.cfg.Modules {
			if err := s.schedule.Regist(&mod{
				Cfg: mc,
				M:   servers[i][j],
			}); err != nil {
				log.ErrorTag(`server`, fmt.Sprintf("服务器 %s 注册失败", mc.Name), err)
				cancel()
				return
			}
		}
	}

	wg.Wait()
	log.InfoTag(`server`, `所有模块初始化成功`)

	for _, s := range serverList {
		s.schedule.Start()
	}

	csignal := make(chan os.Signal)
	signal.Notify(csignal, os.Interrupt, os.Kill)
	<-csignal

	log.InfoTag(`server`, `正在关闭服务器`)

	cancel()

	for _, s := range serverList {
		s.schedule.Stop()
	}

	wg.Wait()
	log.InfoTag(`server`, `服务器关闭`)
}

type server struct {
	schedule *scheduler
	cfg      Conifg
}
