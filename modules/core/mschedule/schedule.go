package mschedule

import (
	"context"
	"j4f/core/log"
	"j4f/core/scheduler"
)

type M_Schedule struct {
}

func (m *M_Schedule) Init(ctx context.Context, name string, cfgPath string) {

}

func (m *M_Schedule) Run() {

}

func (m *M_Schedule) Exec() {
	log.Log()
	scheduler.Exec()
}
