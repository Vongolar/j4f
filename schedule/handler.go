package schedule

import (
	"j4f/module"
	"j4f/task"
)

func (m *MSchedule) GetHandlers() map[int]func(*task.Task) {
	return map[int]func(*task.Task){}
}

func (m *MSchedule) registModule(task task.Task) {
	data := task.Data.(module.IModule)
	if data == nil {
		return
	}

	hs := data.GetHandlers()
	for commandID, handler := range hs {
		if 
	}
}
