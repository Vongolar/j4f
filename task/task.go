package jtask

import (
	"JFFun/data/Derror"
)

//Task 协程间
type Task struct {
	PlayerID string
	Request
	Data interface{}
	Raw  []byte
}

//OK 成功
func (task *Task) OK() {
	task.Error(Derror.Error_ok)
}

//Error 错误
func (task *Task) Error(err Derror.Error) {
	task.Reply(err, nil)
}

//Reply 回应
func (task *Task) Reply(err Derror.Error, data interface{}) {
	if task.Request != nil {
		task.Request.Reply(err, data)
	}
}
