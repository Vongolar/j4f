package schedule

import (
	"j4f/command"
	"j4f/task"
)

func (m *Module) scheduleTaskHandler(t *task.Task) {
	data, exist := t.Data.(*task.Task)
	if !exist {
		t.Reply(nil, command.Err_TASK_DATA_INVAILD)
	}
}
