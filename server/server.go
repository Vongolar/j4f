/*
 * @Author: Vongola
 * @FilePath: \JFFun\server\server.go
 * @Date: 2021-01-23 23:32:09
 * @Description: file content
 * @描述: 文件描述
 * @LastEditTime: 2021-01-29 15:36:34
 * @LastEditors: Vongola
 */

package server

import (
	"JFFun/config"
	"JFFun/jlog"
	"JFFun/module"
	"JFFun/task"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func Run(mods map[string]module.Module) {
	RunServers(mods)
}

func RunServers(servers ...map[string]module.Module) {
	release, cfgPaths := parseFlag()

	if len(servers) > len(cfgPaths) {
		jlog.ErrorWithTag(`server`, `服务器配置不足。`)
		return
	}

	if release && len(servers) > 0 {
		jlog.WarningWithTag(`server`, `发布模式建议将多个模块集成到单个服务器。`)
	}

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i, s := range servers {
		ser := new(server)
		ser.run(ctx, &wg, cfgPaths[i], s)
	}

	sc := make(chan os.Signal)
	signal.Notify(sc, os.Interrupt, os.Kill)
	<-sc
	cancel()
	wg.Wait()
}

type server struct {
	local *localSchedule

	cfg serverConfig
}

type mod struct {
	name string
	c    chan *task.Tuple
	m    module.Module
	cfg  moduleConfig
}

func (s *server) run(ctx context.Context, wg *sync.WaitGroup, cfgPath string, mods map[string]module.Module) {
	err := config.ParseLocalConfig(cfgPath, &s.cfg)
	if err != nil {
		jlog.ErrorWithTag(`server`, fmt.Sprintf("配置文件%s解析错误。", cfgPath))
		return
	}

	if len(s.cfg.Modules) < len(mods) {
		jlog.ErrorWithTag(`server`, fmt.Sprintf("服务器%s 模块配置数不足。", s.cfg.Name))
		return
	}

	s.local = new(localSchedule)

	for name, m := range mods {
		path, ok := s.cfg.Modules[name]
		if !ok {
			jlog.ErrorWithTag(`server`, fmt.Sprintf("服务器%s 模块%s 缺少配置。", s.cfg.Name, name))
			return
		}

		mcfg := new(moduleConfig)
		err := config.ParseLocalConfig(path, mcfg)
		if err != nil {
			jlog.ErrorWithTag(`server`, fmt.Sprintf("服务器%s 模块%s 配置文件解析错误。", s.cfg.Name, name), err)
			return
		}

		err = m.Init(ctx, s.local, name, path)
		if err != nil {
			jlog.ErrorWithTag(`server`, fmt.Sprintf("服务器%s 模块%s 初始化错误。", s.cfg.Name, name), err)
			return
		}

		newModule := &mod{
			name: name,
			c:    make(chan *task.Tuple, mcfg.Buffer),
			m:    m,
			cfg:  *mcfg,
		}
		s.local.registHandlers(newModule)
	}

	for _, m := range s.local.mods {
		s.goRunModule(wg, m)
	}
}

func (s *server) goRunModule(wg *sync.WaitGroup, m *mod) {
	wg.Add(1)
	go func() {
		defer func() {
			err := recover()
			if err == nil {
				wg.Done()
				return
			}
			jlog.ErrorWithTag(`server`, fmt.Sprintf("服务器%s 模块%s 崩溃。", s.cfg.Name, m.name), err)
			if m.cfg.AutoRestart {
				jlog.InfoWithTag(`server`, fmt.Sprintf("服务器%s 模块%s 准备重启。", s.cfg.Name, m.name))
				s.goRunModule(wg, m)
			}
			wg.Done()
		}()
		m.m.Run(m.c)
	}()
}
