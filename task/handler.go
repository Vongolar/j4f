package task

type Handler func(*Task)

func CombineTaskAndHandler(t *Task, h Handler) *TaskHandlerTuple {
	return &TaskHandlerTuple{
		handler: h,
		task:    t,
	}
}

type TaskHandlerTuple struct {
	handler Handler
	task    *Task
}

func (t *TaskHandlerTuple) Exec() {
	t.handler(t.task)
}
