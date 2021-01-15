/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-15 11:23:45
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\module\module.go
 * @Date: 2021-01-15 10:14:10
 * @描述: 文件描述
 */
package module

import (
	"context"
)

type Module interface {
	Init(name string, configPath string) error
	Run(ctx context.Context)
	HandleRunMsg(msg interface{})
}
