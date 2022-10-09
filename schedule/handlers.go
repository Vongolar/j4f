package schedule

import (
	"j4f/module"
	"j4f/task"
	"reflect"
)

type moduleHandlerTuple struct {
	module  module.IModule
	handler func(*task.Task)
}

func (t *moduleHandlerTuple) getModuleType() reflect.Type {
	return reflect.TypeOf(t.module)
}
