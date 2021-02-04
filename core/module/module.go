/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-04 18:16:12
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\module\module.go
 * @Date: 2021-02-04 11:31:39
 * @描述: 文件描述
 */
package module

import (
	"context"
	"j4f/core/task"
)

type Module interface {
	Init(ctx context.Context, name string, cfgPath string) error
	Run(c chan *task.TaskHandleTuple)
}
