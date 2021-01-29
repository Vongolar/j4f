/*
 * @Author: Vongola
 * @FilePath: \JFFun\module\template\template.go
 * @Date: 2021-01-24 21:15:46
 * @Description: file content
 * @描述: 文件描述
 * @LastEditTime: 2021-01-29 15:38:39
 * @LastEditors: Vongola
 */
package template

import (
	"JFFun/data"
	"JFFun/schedule"
	"JFFun/task"
	"context"
)

type MTemplate struct {
	ctx      context.Context
	handlers map[data.Command]task.Handler
}

func (m *MTemplate) Init(ctx context.Context, schedule schedule.Schedule, name string, configPath string) error {
	m.ctx = ctx
	return nil
}

func (m *MTemplate) Run(taskChannel chan *task.Tuple) {
	if len(m.GetHandlers()) == 0 {
		return
	}

	for {
		select {
		case t := <-taskChannel:
			t.Exec()
		case <-m.ctx.Done():
			return
		}
	}
}

func (m *MTemplate) GetHandlers() map[data.Command]task.Handler {
	return m.handlers
}
