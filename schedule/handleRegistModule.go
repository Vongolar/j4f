package schedule

import (
	"context"
	"j4f/command"
	"j4f/module"
	"j4f/task"
	"log"
)

type RegistModuleData struct {
	Module   module.IModule
	Capacity int
	Ctx      context.Context
}

func (m *Module) registModuleHandler(t *task.Task) {
	data, exist := t.Data.(*RegistModuleData)
	if !exist {
		t.Reply(nil, command.Err_TASK_DATA_INVAILD)
	}

	m.wg.Add(1)
	mod := module.CreateModule(data.Module, data.Capacity, data.Ctx)
	m.registHandler(mod)

	go func() {
		defer func() {
			m.wg.Done()
			if err := recover(); err != nil {
				// TODO: 捕获异常，重启模块
				log.Println(err)
				Exec(&task.Task{
					CommandID: command.CMD_REGIST_MODULE,
					Data:      t.Data,
				})
			}
		}()
		mod.Run()
	}()

	t.ReplyOK(nil)
}

func (m *Module) registHandler(mod *module.Module) error {
	if m.handlers == nil {
		m.handlers = make(map[int][]*module.Module)
	}
	cmds := mod.GetCommands()
	for _, cmd := range cmds {
		_, exist := m.handlers[cmd]
		if !exist {
			m.handlers[cmd] = make([]*module.Module, 0)
		}

		if len(m.handlers[cmd]) > 0 { // 有多个模块注册同一个cmd，必须是相同类型
			for _, existMod := range m.handlers[cmd] {
				if existMod == mod {
					log.Println("重复注册")
					continue
				}
				if existMod.GetType() != mod.GetType() {
					return command.Err_REGIST_CONFLICT_TASK_COMMAND
				}
			}
		}

		// 注册成功
		m.handlers[cmd] = append(m.handlers[cmd], mod)
	}
	return nil
}
