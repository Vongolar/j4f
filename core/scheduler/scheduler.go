package scheduler

import (
	"j4f/core/module"
	"j4f/core/task"
)

type ISchedule interface {
	module.Module
	RegistModule(ModuleWithCfg) error
	RunModules() error

	Handle(*task.Task) error
}

type ModuleWithCfg struct {
	Mod    module.Module
	Name   string
	Config string
}
