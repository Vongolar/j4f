package accountMgr

import (
	"context"
	"j4f/core/task"
	"j4f/data/command"
)

type M_AccountMgr struct {
	handlers map[command.Command]task.TaskHandler
}

func (m *M_AccountMgr) Init(ctx context.Context, name string, cfgPath string) error {
	return nil
}

func (m *M_AccountMgr) Run(c chan *task.Task) {
LOOP:
	for {
		select {
		case t := <-c:
			if t == nil {
				break LOOP
			}
			m.GetHandlers()[t.CMD](t)
		}
	}
}
