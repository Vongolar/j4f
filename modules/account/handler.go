/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 14:37:48
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\modules\account\handler.go
 * @Date: 2021-02-07 19:19:58
 * @描述: 文件描述
 */
package account

import (
	"j4f/core/task"
	"j4f/data"
)

func (m *M_Login) GetHandlers() map[data.Command]task.Handler {
	return map[data.Command]task.Handler{}
}

func (m *M_Login) login(task *task.Task) {

}
