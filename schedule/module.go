package schedule

import (
	"j4f/command"
	"j4f/module"
	"j4f/task"
	"log"
	"sync"
)

func Exec(t *task.Task) {
	if execFunc == nil {
		// TODO: 过早调用，错误处理
		return
	}

	scheduleTask := &task.Task{
		CommandID: command.CMD_SCHEDULE_TASK,
		Data:      t,
	}

	execFunc(scheduleTask)
}

func SetExecFunc(f func(t *task.Task)) {
	execFunc = f
}

var execFunc func(t *task.Task)
var scheduler *module.Module

func SetScheduler(s *module.Module) {
	scheduler = s
}

type Module struct {
	wg       *sync.WaitGroup
	handlers map[int][]*module.Module
}

func (m *Module) OnStart() {
	log.Println("调度模块打开")
	m.wg = new(sync.WaitGroup)
}

func (m *Module) OnClose() {
	m.wg.Wait()
	log.Println("调度模块关闭")
}
