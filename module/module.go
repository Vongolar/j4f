package module

import (
	"context"
	"sync"
)

type IModule interface {
	Run(context.Context, *sync.WaitGroup) error
}
