/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-04 18:52:46
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\schedule\schedule.go
 * @Date: 2021-02-04 14:48:25
 * @描述: 文件描述
 */
package schedule

import (
	"context"
	"fmt"
	"j4f/core/log"
	"j4f/core/module"
	"j4f/core/task"
	"sync"
)

type Scheduler struct {
	ctx  context.Context
	wg   *sync.WaitGroup
	mods []*mod
}

func NewSchedule(ctx context.Context, wg *sync.WaitGroup) *Scheduler {
	return &Scheduler{
		ctx: ctx,
		wg:  wg,
	}
}

func (s *Scheduler) Regist(cfg ModuleConfig, m module.Module) error {
	s.wg.Add(1)

	go func() {
		defer s.wg.Done()
		if err := m.Init(s.ctx, cfg.Name, cfg.Config); err != nil {
			log.ErrorTag(`schedule`, fmt.Sprintf("模块 %s 初始化错误", cfg.Name), err)
			panic(err)
		}

		s.mods = append(s.mods, &mod{
			m:   m,
			c:   make(chan *task.TaskHandleTuple, cfg.Buffer),
			cfg: cfg,
		})

		log.InfoTag(`schedule`, fmt.Sprintf("模块 %s 初始化成功", cfg.Name))
	}()

	return nil
}

func (s *Scheduler) Start() {
	for _, m := range s.mods {
		s.goRunMod(m)
	}
}

func (s *Scheduler) Stop() {
	for _, m := range s.mods {
		close(m.c)
	}
}

func (s *Scheduler) goRunMod(m *mod) {
	s.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.ErrorTag(`schedule`, fmt.Sprintf("模块 %s 异常关闭", m.cfg.Name), err)
				log.WarnTag(`schedule`, fmt.Sprintf("模块 %s 重启", m.cfg.Name))
				s.goRunMod(m)
			} else {
				log.InfoTag(`schedule`, fmt.Sprintf("模块 %s 关闭", m.cfg.Name))
			}
			s.wg.Done()
		}()

		m.m.Run(m.c)
	}()
}

type mod struct {
	m   module.Module
	c   chan *task.TaskHandleTuple
	cfg ModuleConfig
}

type ModuleConfig struct {
	Name   string `toml:"name" yaml:"name" json:"name"`
	Config string `toml:"config" yaml:"config" json:"config"`
	Buffer int    `toml:"buffer" yaml:"buffer" json:"buffer"`
}
