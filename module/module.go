package jmodule

import (
	"JFFun/data/Dcommand"
	jtask "JFFun/task"
	"context"
)

//Module 模块
type Module interface {
	Init(cfg string) error
	Run(ctx context.Context, name string)
	GetHandler() map[Dcommand.Command]func(*jtask.Task)
}
