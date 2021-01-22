/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-22 18:29:08
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\gate\handler.go
 * @Date: 2021-01-15 10:14:10
 * @描述: 文件描述
 */
package gate

import (
	"JFFun/data"
	"JFFun/task"
	"fmt"
)

func (m *M_Gate) GetHandlers() map[data.Command]task.Handler {
	return map[data.Command]task.Handler{
		data.Command_ping: m.Ping,
	}
}

func (m *M_Gate) Ping(task *task.Task) {
	fmt.Println("ping")
}
