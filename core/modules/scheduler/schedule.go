package mscheduler

import (
	"fmt"
	"j4f/core/module"
	"j4f/core/task"
	"j4f/data/command"
)

var ErrRejectTask = fmt.Errorf(`服务器开始关闭，拒绝服务`)

func (m *M_Scheduler) RegistModule(mod module.Module) error {
	if m.isClose() {
		return ErrRejectTask
	}
	m.c <- &task.Task{
		CMD:  command.Command_registModule,
		Data: mod,
	}
	return nil
}

func (m *M_Scheduler) RunModules() error {
	if m.isClose() {
		return ErrRejectTask
	}
	m.c <- &task.Task{
		CMD: command.Command_runModules,
	}
	return nil
}
