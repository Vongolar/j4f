/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-15 11:58:29
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
)

func (m *M_Gate) GetHandler() map[data.Command]func(t *task.Task) {
	return map[data.Command]func(t *task.Task){}
}

func (m *M_Gate) HandleRunMsg(msg interface{}) {

}
