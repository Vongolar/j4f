package module

import (
	"JFFun/data/command"
	"JFFun/task"
	"context"
)

//Module 模块
type Module interface {
	GetName() string
	Init() error
	GetHandlers() map[command.Command]func(task *task.Task)
	Run(context.Context)
}
