package module

import (
	"context"
)

type Module interface {
	GetName() string
	Init() error
	Run(context.Context)
}
