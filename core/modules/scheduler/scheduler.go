package mscheduler

import (
	"context"
	"fmt"
	"j4f/core/config"
	moduleconfig "j4f/core/module/config"
	"j4f/core/server"
	"j4f/core/task"
	"j4f/data/command"
	"sync/atomic"
)

type M_Scheduler struct {
	ctx       context.Context
	name      string
	commonCfg moduleconfig.ModuleConfig

	c         chan *task.Task
	closeSign int64
}

func (m *M_Scheduler) Init(ctx context.Context, name string, cfgPath string) error {
	if err := config.ParseFile(cfgPath, &m.commonCfg); err != nil {
		return err
	}

	m.ctx = ctx
	m.name = name

	m.c = make(chan *task.Task, m.commonCfg.Buffer)
	atomic.StoreInt64(&m.closeSign, 0)

	return nil
}

func (m *M_Scheduler) Run() {
	handlers := m.GetHandlers()

LOOP:
	for {
		select {
		case <-m.ctx.Done():
			m.close()
		case t := <-m.c:
			if t == nil {
				break LOOP
			}
			handler, exist := handlers[t.CMD]
			if !exist {
				server.ErrTag(m.name, fmt.Sprintf("No find handler for command : %s .", command.Command_name[int32(t.CMD)]))
				continue
			}
			handler(t)
		}
	}
}

func (m *M_Scheduler) isClose() bool {
	return atomic.LoadInt64(&m.closeSign) > 0
}

func (m *M_Scheduler) close() {
	atomic.AddInt64(&m.closeSign, 1)
	close(m.c)
}
