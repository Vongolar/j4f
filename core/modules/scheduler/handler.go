package mscheduler

import (
	"fmt"
	"j4f/core/config"
	moduleconfig "j4f/core/module/config"
	"j4f/core/scheduler"
	"j4f/core/server"
	"j4f/core/task"
	"j4f/data/command"
	"j4f/data/common"
	"j4f/data/errCode"
)

func (m *M_Scheduler) GetHandlers() map[command.Command]task.TaskHandler {
	if m.handlers == nil {
		m.handlers = map[command.Command]task.TaskHandler{
			command.Command_registModule: m.registModule,
			command.Command_runModules:   m.runModules,
			command.Command_handle:       m.handle,

			command.Command_closeModule:   m.closeModule,
			command.Command_restartModule: m.restartModule,
		}
	}
	return m.handlers
}

func (m *M_Scheduler) registModule(t *task.Task) {
	mc, ok := t.Data.(scheduler.ModuleWithCfg)
	if !ok {
		t.Error(errCode.Code_typeNoMatch)
		return
	}

	for _, moded := range m.modules {
		if moded.m == mc.Mod {
			server.ErrTag(m.name, fmt.Sprintf("模块 %s 重复注册", mc.Name))
			t.Error(errCode.Code_moduleRegistErr)
			return
		}

		if moded.name == mc.Name {
			server.ErrTag(m.name, fmt.Sprintf("模块名 %s 已被占用", mc.Name))
			t.Error(errCode.Code_moduleRegistErr)
			return
		}
	}

	err := mc.Mod.Init(m.ctx, mc.Name, mc.Config)
	if err != nil {
		server.ErrTag(m.name, fmt.Sprintf("模块 %s 初始化错误", mc.Name), err)
		t.Reply(errCode.Code_moduleRegistErr, err)
		return
	}

	var cfg moduleconfig.ModuleConfig
	err = config.ParseFile(mc.Config, &cfg)
	if err != nil {
		server.ErrTag(m.name, fmt.Sprintf("模块 %s 配置文件 %s 解析错误", mc.Name, mc.Config), err)
		t.Reply(errCode.Code_moduleRegistErr, err)
		return
	}

	channel := make(chan *task.Task, cfg.Buffer)
	newModule := &mod{name: mc.Name, cftPath: mc.Config, m: mc.Mod, c: channel, cfg: cfg, handlers: mc.Mod.GetHandlers()}

	//注册方法
	for cmd := range newModule.handlers {
		m.handlerMap[cmd] = append(m.handlerMap[cmd], newModule)
	}

	//性能缓存
	if m.cfg.Profile > 0 {
		newModule.profileQueue = createFixIntergerQueue(m.cfg.Profile)
		newModule.errProfileQueue = createFixBooleanQueue(m.cfg.Profile)
	}

	m.modules = append(m.modules, newModule)
	server.InfoTag(m.name, fmt.Sprintf("模块 %s 注册完成", mc.Name))
	t.Ok()
}

func (m *M_Scheduler) runModules(t *task.Task) {
	for _, mod := range m.modules {
		m.goRunModule(mod)
	}
}

func (m *M_Scheduler) closeModule(t *task.Task) {
	name := t.Data.(*common.ModuleName)
	m.closeM(name.GetName())
	t.Ok()
}

func (m *M_Scheduler) restartModule(t *task.Task) {
	name := t.Data.(*common.ModuleName)
	m.restartM(name.GetName())
	t.Ok()
}

func (m *M_Scheduler) goRunModule(mod *mod) {
	mod.Lock()
	mod.enable = true
	mod.Unlock()

	m.wg.Add(1)
	go func() {
		defer func() {
			mod.Lock()
			mod.enable = false
			mod.Unlock()

			if err := recover(); err != nil {
				server.ErrTag(mod.name, `模块异常退出`, err)
				if mod.cfg.AutoRestart {
					server.InfoTag(mod.name, `自动重启`)
					m.goRunModule(mod)
				}
			} else {
				server.InfoTag(mod.name, `模块安全退出`)
			}
			m.wg.Done()
		}()
		mod.m.Run(mod.c)
	}()
}

func (m *M_Scheduler) handle(t *task.Task) {
	ct := t.Data.(*task.Task)
	mods := m.getEnableModulesByCMD(ct.CMD)
	if len(mods) == 0 {
		ct.Error(errCode.Code_noFindModuleForTask)
		return
	}

	mods[0].c <- ct
}

func (m *M_Scheduler) closeM(name string) {
	cm := m.findModule(name)
	if cm == nil {
		return
	}

	cm.RLock()
	if !cm.enable {
		cm.RUnlock()
		return
	}
	cm.RUnlock()
	close(cm.c)

	for cm.enable {
	}
}

func (m *M_Scheduler) restartM(name string) {
	m.closeM(name)
	mod := m.findModule(name)

	if mod == nil {
		return
	}

	mod.c = make(chan *task.Task, mod.cfg.Buffer)

	//性能缓存
	if m.cfg.Profile > 0 {
		mod.profileQueue = createFixIntergerQueue(m.cfg.Profile)
		mod.errProfileQueue = createFixBooleanQueue(m.cfg.Profile)
	}

	mod.m.Init(m.ctx, mod.name, mod.cftPath)
	m.goRunModule(mod)
}
