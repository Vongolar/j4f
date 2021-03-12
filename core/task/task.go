package task

import (
	"j4f/data/command"
	"j4f/data/errCode"
)

type Task struct {
	CMD  command.Command
	Data interface{}
}

type TaskHandler func(*Task)

func (t *Task) Reply(code errCode.Code, ext interface{}) {

}

func (t *Task) Ok() {
	t.Error(errCode.Code_ok)
}

func (t *Task) Error(code errCode.Code) {
	t.Reply(code, nil)
}
