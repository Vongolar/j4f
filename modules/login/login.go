/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-07 19:01:56
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\modules\login\login.go
 * @Date: 2021-02-07 18:54:49
 * @描述: 文件描述
 */

package login

import (
	"context"
	"j4f/core/scheduler"
	"j4f/core/task"
)

type M_Login struct {
	name      string
	ctx       context.Context
	scheduler scheduler.Scheduler
}

func (m *M_Login) Init(ctx context.Context, name string, cfgPath string) error {
	m.name = name
	m.ctx = ctx

	return nil
}

func (m *M_Login) Run(c chan *task.TaskHandleTuple, s scheduler.Scheduler) {
	m.scheduler = s

LOOP:
	for {
		select {
		case t := <-c:
			if t == nil {
				break LOOP
			}
		}
	}
}
