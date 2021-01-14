package module

import (
	"context"
)

type Module interface {
	Init(name string, configPath string) error
	Run(ctx context.Context)
}
