package gate

import (
	"context"
)

type M_Gate struct {
}

func (m *M_Gate) Init(name string, configPath string) error {
	return nil
}

func (m *M_Gate) Run(ctx context.Context) {

}
