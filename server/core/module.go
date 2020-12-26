/*
 * @Author: Vongola
 * @FilePath: /JFFun/server/core/module.go
 * @Date: 2020-12-26 16:03:41
 * @Description: Module Interface
 * @描述: 模块接口
 * @LastEditTime: 2020-12-26 16:10:27
 * @LastEditors: Vongola
 */

package core

import "context"

type Module interface {
	Init(cfg string) error
	Run(ctx context.Context)
}
