/*
 * @Author: Vongola
 * @FilePath: /JFFun/task/task.go
 * @Date: 2021-01-24 00:02:03
 * @Description: file content
 * @描述: 文件描述
 * @LastEditTime: 2021-01-24 22:31:24
 * @LastEditors: Vongola
 */

package task

import (
	"JFFun/data"
)

type Task struct {
	CommandID data.Command
}

type Handler func(*Task)

func NewHandleTuple(t *Task, h Handler) *Tuple {
	return &Tuple{
		task:    t,
		handler: h,
	}
}

type Tuple struct {
	task    *Task
	handler Handler
}

func (t *Tuple) Exec() {
	t.handler(t.task)
}
