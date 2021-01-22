/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-22 18:56:14
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\schedule\schedule.go
 * @Date: 2021-01-15 10:14:10
 * @描述: 文件描述
 */
package schedule

import (
	"JFFun/config"
	"JFFun/jlog"
	"JFFun/module"
	"JFFun/task"
	"context"
	"fmt"
	"sync"
)

type Schedule struct {
	modules map[string]*mod
}

type mod struct {
	m      module.Module
	buffer chan task.TaskHandlerTuple
	cfg    moduleConfig
}

func (s *Schedule) RegistModule(name string, m module.Module, cfgPath string) error {
	cfg := new(moduleConfig)
	if err := config.ParseLocalConfig(cfgPath, cfg); err != nil {
		return err
	}

	if s.modules == nil {
		s.modules = make(map[string]*mod)
	}

	s.modules[name] = &mod{
		m:      m,
		cfg:    *cfg,
		buffer: make(chan task.TaskHandlerTuple, cfg.Buffer),
	}

	return nil
}

func (s *Schedule) Run(ctx context.Context, wg *sync.WaitGroup) {
	for k, m := range s.modules {
		s.goRunModule(ctx, wg, k, m)
	}
}

func (s *Schedule) goRunModule(ctx context.Context, wg *sync.WaitGroup, name string, mod *mod) {
	wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				jlog.Error(fmt.Sprintf("module %s panic : %v", name, err))

				if mod.cfg.AutoRestart {
					jlog.Info(fmt.Sprintf("module %s auto restart", name))
					s.goRunModule(ctx, wg, name, mod)
				}
			}
			wg.Done()
		}()
		mod.m.Run(ctx, s, mod.buffer)
	}()
}
