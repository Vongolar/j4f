/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-19 11:57:25
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\module\module.go
 * @Date: 2021-02-19 11:56:56
 * @描述: 文件描述
 */
package module

import (
	"context"
	"j4f/core/task"
	"j4f/data/command"
)

type Module interface {
	Init(ctx context.Context, name string, cfgPath string) error
	Run(chan *task.Task)
	GetHandlers() map[command.Command]task.TaskHandler
}
