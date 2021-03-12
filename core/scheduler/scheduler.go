package scheduler

import (
	"j4f/core/module"
)

type ISchedule interface {
	module.Module
	RegistModule(ModuleWithCfg) error
	RunModules() error
}

type ModuleWithCfg struct {
	Mod    module.Module
	Name   string
	Config string
}
