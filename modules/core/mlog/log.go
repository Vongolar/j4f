/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-19 11:59:05
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\modules\mlog\log.go
 * @Date: 2021-02-19 11:58:52
 * @描述: 文件描述
 */

package mlog

import (
	"context"
	"j4f/core/log"
	"j4f/core/scheduler"
)

type M_Log struct {
}

func (m *M_Log) Init(ctx context.Context, name string, cfgPath string) {

}

func (m *M_Log) Run() {

}

func (m *M_Log) Log() {
	log.Log()
	scheduler.Exec()
}
