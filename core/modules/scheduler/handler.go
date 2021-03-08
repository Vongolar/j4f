package mscheduler

import (
	"j4f/core/task"
	"j4f/data/command"
)

func (m *M_Scheduler) GetHandlers() map[command.Command]task.TaskHandler {
	return map[command.Command]task.TaskHandler{
		command.Command_registModule: m.registModule,
	}
}

func (m *M_Scheduler) registModule(t *task.Task) {

}
