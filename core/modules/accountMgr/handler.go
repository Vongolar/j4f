package accountMgr

import (
	"j4f/core/task"
	"j4f/data/command"
)

func (m *M_AccountMgr) GetHandlers() map[command.Command]task.TaskHandler {
	if m.handlers == nil {
		m.handlers = map[command.Command]task.TaskHandler{}
	}
	return m.handlers
}
