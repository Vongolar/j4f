/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 18:22:56
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\server\schedule.go
 * @Date: 2021-02-08 15:32:07
 * @描述: 文件描述
 */

package server

import (
	"j4f/core/task"
	"j4f/data"
	"math"
)

func (s *scheduler) Exec(t *task.Task) {
	handlers, ok := s.handlers[t.CMD]
	if !ok || len(handlers) == 0 {
		t.Err(data.Error_noSupportCommand)
		return
	}

	min := math.MaxInt64
	var execHandler *handler
	for _, h := range handlers {
		execing := len(h.m.C)
		if execing < min {
			min = execing
			execHandler = h
		}
	}
	if execHandler == nil {
		t.Err(data.Error_noSupportCommand)
		return
	}

	execHandler.m.C <- task.NewTaskHandlerTuple(t, execHandler.handler)
}

func (s *scheduler) ExecLocal(t *task.Task) {
	handlers, ok := s.handlers[t.CMD]
	if !ok || len(handlers) == 0 {
		t.Err(data.Error_noSupportCommand)
		return
	}

	min := math.MaxInt64
	var execHandler *handler
	for _, h := range handlers {
		execing := len(h.m.C)
		if execing < min {
			min = execing
			execHandler = h
		}
	}
	if execHandler == nil {
		t.Err(data.Error_noSupportCommand)
		return
	}

	execHandler.m.C <- task.NewTaskHandlerTuple(t, execHandler.handler)
}

func (s *scheduler) ExecMutli(task *task.Task, mutli int) {

}
