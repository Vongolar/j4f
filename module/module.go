/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-22 19:03:15
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\module\module.go
 * @Date: 2021-01-15 10:14:10
 * @描述: 文件描述
 */
package module

import (
	"JFFun/data"
	"JFFun/task"
	"context"
)

type Module interface {
	Init(serverName string, configPath string) error
	Run(context.Context, task.Execer, chan task.TaskHandlerTuple)
	GetHandlers() map[data.Command]task.Handler
}
