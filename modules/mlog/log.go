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
func (m *M_Log) Run() {

}
func (m *M_Log) GetHandlers() map[command.Command]task.TaskHandler {
	return map[command.Command]task.TaskHandler{}
}
