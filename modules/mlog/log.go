package mlog

import (
	"context"
	"j4f/core/task"
	"j4f/data/command"
)

type M_Log struct {
}

func (m *M_Log) Init(ctx context.Context, name string, cfgPath string) error {
	return nil
}

var first = true

func (m *M_Log) Run(c chan *task.Task) {
	if first {
		first = false
		panic("log panic")
	}
LOOP:
	for {
		select {
		case t := <-c:
			if t == nil {
				break LOOP
			}
		}
	}
}

func (m *M_Log) GetHandlers() map[command.Command]task.TaskHandler {
	return map[command.Command]task.TaskHandler{}
}
