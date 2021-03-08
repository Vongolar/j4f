package schduler

import (
	"j4f/core/module"
)

type ISchedule interface {
	module.Module
	RegistModule(module.Module) error
	RunModules() error
}
