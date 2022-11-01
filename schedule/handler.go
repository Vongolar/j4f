package schedule

import (
	"j4f/command"
	"j4f/module"
)

func (m *Module) GetHandlers() map[int]module.TaskHandle {
	return map[int]module.TaskHandle{
		command.CMD_REGIST_MODULE: m.registModuleHandler,
	}
}
