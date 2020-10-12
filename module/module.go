package module

import (
	Jcommand "JFFun/data/command"
	"JFFun/task"
	"context"
)

type Module interface {
	GetName() string
	Init() error
	GetHandlers() map[Jcommand.Command]func(task *task.Task)
	Run(context.Context)
}
