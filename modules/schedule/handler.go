/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 18:08:18
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\modules\schedule\handler.go
 * @Date: 2021-02-08 14:38:09
 * @描述: 文件描述
 */
package schedule

import (
	"j4f/core/task"
	"j4f/data"
)

func (m *M_Schedule) GetHandlers() map[data.Command]task.Handler {
	return map[data.Command]task.Handler{
		data.Command_ping: m.Ping,
	}
}

func (m *M_Schedule) Ping(task *task.Task) {
	task.OK(nil)
}
