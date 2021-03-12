package mscheduler

import (
	"fmt"
	"j4f/core/config"
	moduleconfig "j4f/core/module/config"
	"j4f/core/scheduler"
	"j4f/core/server"
	"j4f/core/task"
	"j4f/data/command"
	"j4f/data/errCode"
)

func (m *M_Scheduler) GetHandlers() map[command.Command]task.TaskHandler {
	return map[command.Command]task.TaskHandler{
		command.Command_registModule: m.registModule,
		command.Command_runModules:   m.runModules,
	}
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
	m.modules = append(m.modules, &mod{name: mc.Name, m: mc.Mod, c: channel, cfg: cfg})
	server.InfoTag(m.name, fmt.Sprintf("模块 %s 注册完成", mc.Name))
}

func (m *M_Scheduler) runModules(t *task.Task) {
	for _, mod := range m.modules {
		m.goRunModule(mod)
	}
}

func (m *M_Scheduler) goRunModule(mod *mod) {
	m.wg.Add(1)
	go func() {
		defer func() {
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
