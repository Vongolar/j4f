/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 18:16:56
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\server\scheduler.go
 * @Date: 2021-02-04 14:48:25
 * @描述: 文件描述
 */
package server

import (
	"context"
	"fmt"
	"j4f/core/task"
	"j4f/data"
	"sync"
)

type scheduler struct {
	name string
	sync.RWMutex
	ctx      context.Context
	wg       *sync.WaitGroup
	mods     []*mod
	handlers map[data.Command][]*handler
}

type handler struct {
	m       *mod
	cmd     data.Command
	handler task.Handler
}

func newSchedule(ctx context.Context, wg *sync.WaitGroup) *scheduler {
	s := &scheduler{
		ctx:      ctx,
		wg:       wg,
		handlers: make(map[data.Command][]*handler),
	}
	return s
}

func (s *scheduler) Regist(m *mod) error {
	s.wg.Add(1)

	go func() {
		defer s.wg.Done()
		if err := m.M.Init(s.ctx, m.Cfg.Name, m.Cfg.Config); err != nil {
			s.ErrorTag(`schedule`, fmt.Sprintf("模块 %s 初始化错误", m.Cfg.Name), err)
			panic(err)
		}

		m.handlers = m.M.GetHandlers()
		for cmd, h := range m.handlers {
			s.handlers[cmd] = append(s.handlers[cmd], &handler{m: m, cmd: cmd, handler: h})
		}

		if len(m.handlers) > 0 {
			m.C = make(chan *task.TaskHandleTuple, m.Cfg.Buffer)
		}
		s.mods = append(s.mods, m)

		s.InfoTag(`schedule`, fmt.Sprintf("模块 %s 初始化成功", m.Cfg.Name))
	}()

	return nil
}

func (s *scheduler) Start() {
	for _, m := range s.mods {
		s.goRunMod(m)
	}
}

func (s *scheduler) Stop() {
	for _, m := range s.mods {
		s.wg.Add(1)
		c := m.C
		wg := s.wg
		name := m.Cfg.Name
		go func() {
			defer func() {
				if err := recover(); err != nil {
					s.WarnTag(`schedule`, fmt.Sprintf("模块 %s 关闭异常，可忽略", name), err)
				}
				wg.Done()
			}()
			close(c)
		}()
	}
}

func (s *scheduler) goRunMod(m *mod) {
	s.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				s.ErrorTag(`schedule`, fmt.Sprintf("模块 %s 异常关闭", m.Cfg.Name), err)
				s.WarnTag(`schedule`, fmt.Sprintf("模块 %s 重启", m.Cfg.Name))
				s.goRunMod(m)
			}
			s.wg.Done()
		}()

		m.M.Run(m.C, s)
	}()
}
