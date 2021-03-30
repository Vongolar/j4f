package server

import "j4f/core/task"

func Handle(task *task.Task) error {
	return scheduleModuler.Handle(task)
}
