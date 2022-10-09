package schedule

import (
	"context"
	"sync"
)

type MSchedule struct {
	handlers map[int][]*moduleHandlerTuple
}

func (m *MSchedule) Run(ctx context.Context, wg *sync.WaitGroup) error {
	go m.run(ctx, wg)
	return nil
}

func (m *MSchedule) run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	select {
	case <-ctx.Done():
		return
	}
}
