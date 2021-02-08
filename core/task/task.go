/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 16:43:53
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\task\task.go
 * @Date: 2021-02-04 17:44:05
 * @描述: 文件描述
 */
package task

import (
	"j4f/core/message"
	"j4f/data"
)

type Task struct {
	CMD   data.Command
	Data  interface{}
	Seria message.Seria
	Request
}

type Handler func(task *Task)

func NewTaskHandlerTuple(task *Task, handler Handler) *TaskHandleTuple {
	return &TaskHandleTuple{task, handler}
}

type TaskHandleTuple struct {
	task    *Task
	handler Handler
}

func (t *TaskHandleTuple) Exec() {
	t.handler(t.task)
}
