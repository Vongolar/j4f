package module

import (
	Jcommand "JFFun/data/command"
	Jtask "JFFun/task"
	"context"
)

//Module 模块接口
type Module interface {
	Init(cfg []byte) error
	GetName() string
	GetHandlers() map[Jcommand.Command]func(task *Jtask.Task)
	Run(context.Context)
}
