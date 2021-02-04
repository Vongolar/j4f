/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-04 19:47:55
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
	"j4f/core/log"
	"j4f/core/task"
	"sync"
)

type scheduler struct {
	ctx  context.Context
	wg   *sync.WaitGroup
	mods []*mod
}

func newSchedule(ctx context.Context, wg *sync.WaitGroup) *scheduler {
	return &scheduler{
		ctx: ctx,
		wg:  wg,
	}
}

func (s *scheduler) Regist(m *mod) error {
	s.wg.Add(1)

	go func() {
		defer s.wg.Done()
		if err := m.M.Init(s.ctx, m.Cfg.Name, m.Cfg.Config); err != nil {
			log.ErrorTag(`schedule`, fmt.Sprintf("模块 %s 初始化错误", m.Cfg.Name), err)
			panic(err)
		}

		m.C = make(chan *task.TaskHandleTuple, m.Cfg.Buffer)
		s.mods = append(s.mods, m)

		log.InfoTag(`schedule`, fmt.Sprintf("模块 %s 初始化成功", m.Cfg.Name))
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
		close(m.C)
	}
}

func (s *scheduler) goRunMod(m *mod) {
	s.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.ErrorTag(`schedule`, fmt.Sprintf("模块 %s 异常关闭", m.Cfg.Name), err)
				log.WarnTag(`schedule`, fmt.Sprintf("模块 %s 重启", m.Cfg.Name))
				s.goRunMod(m)
			} else {
				log.InfoTag(`schedule`, fmt.Sprintf("模块 %s 关闭", m.Cfg.Name))
			}
			s.wg.Done()
		}()

		m.M.Run(m.C, s)
	}()
}
