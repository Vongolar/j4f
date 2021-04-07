package task

import (
	"j4f/core/request"
	"j4f/data/command"
	"j4f/data/errCode"
	"j4f/define"
)

type Task struct {
	CMD     command.Command
	Author  define.Auth
	Data    interface{}
	Request request.Request
}

type TaskHandler func(*Task)

func (t *Task) Reply(code errCode.Code, ext interface{}) {
	if t.Request != nil {
		t.Request.Reply(code, ext)
	}
}

func (t *Task) Ok() {
	t.Error(errCode.Code_ok)
}

func (t *Task) Error(code errCode.Code) {
	t.Reply(code, nil)
}
