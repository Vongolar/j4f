/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-05 12:07:05
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\scheduler\scheduler.go
 * @Date: 2021-02-04 19:37:46
 * @描述: 文件描述
 */

package scheduler

type Scheduler interface {
	Info(a ...interface{})
	InfoTag(tag string, a ...interface{})
	Warn(a ...interface{})
	WarnTag(tag string, a ...interface{})
	Error(a ...interface{})
	ErrorTag(tag string, a ...interface{})
}
