/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-22 19:03:06
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\gate\gate.go
 * @Date: 2021-01-15 10:14:10
 * @描述: 文件描述
 */
package gate

import (
	"JFFun/task"
	"context"
	"time"
)

type M_Gate struct {
	execer task.Execer
}

func (m *M_Gate) Init(name string, configPath string) error {
	return nil
}

func (m *M_Gate) Run(ctx context.Context, execer task.Execer, buffer chan task.TaskHandlerTuple) {
	m.execer = execer
	for {
		select {
		case <-ctx.Done():
			return
		case tuple := <-buffer:
			tuple.Exec()
		case <-time.Tick(time.Second):
			// execer.ExecuteLocal()
		}
	}
}
