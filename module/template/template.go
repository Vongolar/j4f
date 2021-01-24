/*
 * @Author: Vongola
 * @FilePath: /JFFun/module/template/template.go
 * @Date: 2021-01-24 21:15:46
 * @Description: file content
 * @描述: 文件描述
 * @LastEditTime: 2021-01-24 22:38:37
 * @LastEditors: Vongola
 */
package template

import (
	"JFFun/data"
	"JFFun/schedule"
	"JFFun/task"
	"context"
	"time"
)

type MTemplate struct {
	ctx context.Context
}

func (m *MTemplate) Init(ctx context.Context, schedule schedule.Schedule, name string, configPath string) error {
	m.ctx = ctx
	return nil
}

func (m *MTemplate) Run(taskChannel chan *task.Tuple) {
	for {
		select {
		case t := <-taskChannel:
			t.Exec()
		case <-m.ctx.Done():
			return
		case <-time.Tick(time.Second * 2):
			panic("I am bug")
		}
	}
}

func (m *MTemplate) GetHandlers() map[data.Command]task.Handler {
	return map[data.Command]task.Handler{}
}
