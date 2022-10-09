package module

import (
	"context"
	"j4f/task"
	"sync"
)

type IModule interface {
	Run(context.Context, *sync.WaitGroup) error
	GetHandlers() map[int]func(*task.Task)
}
