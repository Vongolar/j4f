/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 17:06:20
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\modules\http\handler.go
 * @Date: 2021-02-08 14:38:09
 * @描述: 文件描述
 */
package http

import (
	"j4f/core/task"
	"j4f/data"
)

func (m *M_Http) GetHandlers() map[data.Command]task.Handler {
	return map[data.Command]task.Handler{}
}
