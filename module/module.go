/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-24 22:37:57
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: /JFFun/module/module.go
 * @Date: 2021-01-15 10:14:10
 * @描述: 文件描述
 */
package module

import (
	"JFFun/data"
	"JFFun/schedule"
	"JFFun/task"
	"context"
)

type Module interface {
	Init(ctx context.Context, schedule schedule.Schedule, name string, configPath string) error
	Run(taskChannel chan *task.Tuple)
	GetHandlers() map[data.Command]task.Handler
}
