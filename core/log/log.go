/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-04 14:39:30
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\log\log.go
 * @Date: 2021-02-04 11:53:01
 * @描述: 文件描述
 */
package log

import "fmt"

func Info(a ...interface{}) {
	fmt.Println(append([]interface{}{`[INFO]`}, a...)...)
}

func InfoTag(tag string, a ...interface{}) {
	Info(append([]interface{}{fmt.Sprintf("[%s]", tag)}, a...)...)
}

func Warn(a ...interface{}) {
	fmt.Println(append([]interface{}{`[WARN]`}, a...)...)
}

func WarnTag(tag string, a ...interface{}) {
	Warn(append([]interface{}{fmt.Sprintf("[%s]", tag)}, a...)...)
}

func Error(a ...interface{}) {
	fmt.Println(append([]interface{}{`[ERRO]`}, a...)...)
}

func ErrorTag(tag string, a ...interface{}) {
	Error(append([]interface{}{fmt.Sprintf("[%s]", tag)}, a...)...)
}
