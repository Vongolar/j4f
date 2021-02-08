/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 17:07:06
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\modules\schedule\schedule.go
 * @Date: 2021-02-08 17:06:32
 * @描述: 文件描述
 */

package schedule

import (
	"context"
	"j4f/core/scheduler"
	"j4f/core/task"
)

type M_Schedule struct {
	name      string
	ctx       context.Context
	scheduler scheduler.Scheduler
}

func (m *M_Schedule) Init(ctx context.Context, name string, cfgPath string) error {
	m.name = name
	m.ctx = ctx

	return nil
}

func (m *M_Schedule) Run(c chan *task.TaskHandleTuple, s scheduler.Scheduler) {
	m.scheduler = s

LOOP:
	for {
		select {
		case t := <-c:
			if t == nil {
				break LOOP
			}
			t.Exec()
		}
	}
}
