package mscheduler

import (
	"fmt"
	"j4f/core/request"
	"j4f/core/scheduler"
	"j4f/core/task"
	"j4f/data/command"
	"j4f/data/errCode"
)

var ErrRejectTask = fmt.Errorf(`服务器开始关闭，拒绝服务`)

func (m *M_Scheduler) RegistModule(mod scheduler.ModuleWithCfg) error {
	if m.isClose() {
		return ErrRejectTask
	}
	r := request.CreateSyncRequest()
	m.c <- &task.Task{
		CMD:     command.Command_registModule,
		Data:    mod,
		Request: r,
	}

	if ec, _ := r.Wait(); ec != errCode.Code_ok {
		return fmt.Errorf("注册模块错误。%s", errCode.Code_name[int32(ec)])
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

func (m *M_Scheduler) Handle(t *task.Task) error {
	if m.isClose() {
		return ErrRejectTask
	}
	m.c <- &task.Task{
		CMD:  command.Command_handle,
		Data: t,
	}
	return nil
}
