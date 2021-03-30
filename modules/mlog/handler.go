package mlog

import (
	"fmt"
	"j4f/core/task"
	"j4f/data/command"
)

func (m *M_Log) GetHandlers() map[command.Command]task.TaskHandler {
	if m.handlers == nil {
		m.handlers = map[command.Command]task.TaskHandler{
			command.Command_log: m.log,
		}
	}
	return m.handlers
}

func (m *M_Log) log(t *task.Task) {
	fmt.Println(t.Data)
}
