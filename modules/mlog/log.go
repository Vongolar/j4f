package mlog

import (
	"context"
	"j4f/core/server"
	"j4f/core/task"
	"j4f/data/command"
)

type M_Log struct {
	handlers map[command.Command]task.TaskHandler
}

func (m *M_Log) Init(ctx context.Context, name string, cfgPath string) error {
	return nil
}

func (m *M_Log) Run(c chan *task.Task) {
	server.CloseLogBuffer()

LOOP:
	for {
		select {
		case t := <-c:
			if t == nil {
				break LOOP
			}
			m.handlers[t.CMD](t)
		}
	}
}
