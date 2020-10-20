package schedule

import (
	Jcommand "JFFun/data/command"
	Jerror "JFFun/data/error"
	Jmodule "JFFun/module"
	Jtask "JFFun/task"
)

type worker struct {
	mod         Jmodule.Module
	taskChannel chan *task
	handlers    map[Jcommand.Command]func(task *Jtask.Task)
	delay       int64 //延迟毫秒
}

func (w *worker) handleTask(task *task) {
	if hanler, ok := w.handlers[task.cmd]; ok {
		hanler(task.task)
		return
	}
	task.task.Error(Jerror.Error_commandNotAllow)
}
