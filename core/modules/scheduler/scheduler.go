package mscheduler

import (
	"context"
	"j4f/core/config"
	"j4f/core/module"
	moduleconfig "j4f/core/module/config"
	"j4f/core/task"
	"j4f/data/command"
	"sync"
	"sync/atomic"
)

type M_Scheduler struct {
	ctx       context.Context
	name      string
	commonCfg moduleconfig.ModuleConfig
	cfg       mconfig

	c         chan *task.Task
	closeSign int64

	wg      sync.WaitGroup
	modules []*mod

	handlerMap map[command.Command][]*mod
}

type mod struct {
	name string
	m    module.Module
	c    chan *task.Task
	cfg  moduleconfig.ModuleConfig

	handlers map[command.Command]task.TaskHandler

	profileQueue    *fixIntergerQueue
	errProfileQueue *fixBoolenQueue
}

func (m *M_Scheduler) Init(ctx context.Context, name string, cfgPath string) error {
	if err := config.ParseFile(cfgPath, &m.commonCfg); err != nil {
		return err
	}

	if err := config.ParseFile(cfgPath, &m.cfg); err != nil {
		return err
	}

	m.ctx = ctx
	m.name = name

	m.c = make(chan *task.Task, m.commonCfg.Buffer)
	m.handlerMap = make(map[command.Command][]*mod)

	atomic.StoreInt64(&m.closeSign, 0)
	return nil
}

func (m *M_Scheduler) Run(chan *task.Task) {
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
				//server.ErrTag(m.name, fmt.Sprintf("No find handler for command : %s .", command.Command_name[int32(t.CMD)]))
				continue
			}

			// subTask, middle := t.Data.(*task.Task)
			// if middle && subTask.CMD > command.Command_innerMax {
			// 	server.InfoTag(m.name, fmt.Sprintf("%s", command.Command_name[int32(t.CMD)]))
			// }
			handler(t)
		}
	}

	for _, mod := range m.modules {
		close(mod.c)
	}
	m.wg.Wait()
}

func (m *M_Scheduler) isClose() bool {
	return atomic.LoadInt64(&m.closeSign) > 0
}

func (m *M_Scheduler) close() {
	if !m.isClose() {
		atomic.AddInt64(&m.closeSign, 1)
		close(m.c)
	}
}
