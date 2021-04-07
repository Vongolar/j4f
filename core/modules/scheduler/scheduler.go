package mscheduler

import (
	"context"
	"j4f/core/config"
	"j4f/core/module"
	moduleconfig "j4f/core/module/config"
	"j4f/core/task"
	"j4f/data/command"
	"j4f/data/errCode"
	"j4f/define"
	"sync"
	"sync/atomic"
)

type M_Scheduler struct {
	ctx       context.Context
	name      string
	commonCfg moduleconfig.ModuleConfig
	cfg       mconfig
	self      *mod

	c         chan *task.Task
	closeSign int64

	wg      sync.WaitGroup
	modules []*mod

	handlerMap map[command.Command][]*mod
	handlers   map[command.Command]task.TaskHandler
}

type mod struct {
	name    string
	m       module.Module
	c       chan *task.Task
	cfg     moduleconfig.ModuleConfig
	cftPath string

	handlers map[command.Command]task.TaskHandler

	sync.RWMutex
	enable          bool
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

	m.self = &mod{m: m, enable: true, c: m.c, handlers: m.GetHandlers(), name: name, cfg: m.commonCfg}
	//性能缓存
	if m.cfg.Profile > 0 {
		m.self.profileQueue = createFixIntergerQueue(m.cfg.Profile)
		m.self.errProfileQueue = createFixBooleanQueue(m.cfg.Profile)
	}

	atomic.StoreInt64(&m.closeSign, 0)
	return nil
}

func (m *M_Scheduler) Run(chan *task.Task) {
LOOP:
	for {
		select {
		case <-m.ctx.Done():
			m.close()
		case t := <-m.c:
			if t == nil {
				break LOOP
			}

			handler, exist := m.GetHandlers()[t.CMD]
			if !exist {
				//server.ErrTag(m.name, fmt.Sprintf("No find handler for command : %s .", command.Command_name[int32(t.CMD)]))
				continue
			}

			subTask, middle := t.Data.(*task.Task)
			if middle && !define.HasAuthority(subTask.Author, subTask.CMD) {
				subTask.Error(errCode.Code_noAuthority)
				continue
			}

			// if middle && subTask.CMD > command.Command_innerMax {
			// 	server.InfoTag(m.name, fmt.Sprintf("%s", command.Command_name[int32(t.CMD)]))
			// }
			handler(t)
		}
	}

	for _, mod := range m.modules {
		mod.RLock()
		if mod.enable {
			close(mod.c)
		}
		mod.RUnlock()
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

func (m *M_Scheduler) findModule(name string) *mod {
	for _, module := range m.modules {
		if module.name == name {
			return module
		}
	}
	return nil
}

func (m *M_Scheduler) getEnableModulesByCMD(cmd command.Command) []*mod {
	mods := m.handlerMap[cmd]
	var res []*mod

	if _, ok := m.GetHandlers()[cmd]; ok {
		res = append(res, m.self)
	}

	for _, mod := range mods {
		mod.RLock()
		if mod.enable {
			res = append(res, mod)
		}
		mod.RUnlock()
	}

	return res
}
