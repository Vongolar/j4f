/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-19 15:49:52
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\log\log.go
 * @Date: 2021-02-19 15:49:51
 * @描述: 1.在shedule和log模块运行之前，打印在标准输出和缓存中 2.运行之后，输出到log模块
 */

package log

import (
	"fmt"
	"log"
	"strings"
	"unsafe"
)

type Logger interface {
	Log()
}

type buffLogWriter struct {
	builder strings.Builder
}

var blw *buffLogWriter

func (w *buffLogWriter) Write(p []byte) (n int, err error) {
	fmt.Print(*(*string)(unsafe.Pointer(&p)))
	return w.builder.Write(p)
}

func Log(a ...interface{}) {
	if blw == nil {
		blw = new(buffLogWriter)
		log.SetOutput(blw)
	}
}
