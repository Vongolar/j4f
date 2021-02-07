/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-07 18:20:18
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\server\log.go
 * @Date: 2021-02-05 12:02:17
 * @描述: 文件描述
 */
package server

import "fmt"

func info(a ...interface{}) {
	fmt.Println(append([]interface{}{`[INFO]`}, a...)...)
}

func infoTag(tag string, a ...interface{}) {
	info(append([]interface{}{fmt.Sprintf("[%s]", tag)}, a...)...)
}

func warn(a ...interface{}) {
	fmt.Println(append([]interface{}{`[WARN]`}, a...)...)
}

func warnTag(tag string, a ...interface{}) {
	warn(append([]interface{}{fmt.Sprintf("[%s]", tag)}, a...)...)
}

func err(a ...interface{}) {
	fmt.Println(append([]interface{}{`[ERRO]`}, a...)...)
}

func errTag(tag string, a ...interface{}) {
	err(append([]interface{}{fmt.Sprintf("[%s]", tag)}, a...)...)
}

func (s *scheduler) Info(a ...interface{}) {
	info(append([]interface{}{fmt.Sprintf("[%s]", s.name)}, a...)...)
}
func (s *scheduler) InfoTag(tag string, a ...interface{}) {
	info(append([]interface{}{fmt.Sprintf("[%s]", s.name), fmt.Sprintf("[%s]", tag)}, a...)...)
}
func (s *scheduler) Warn(a ...interface{}) {
	warn(append([]interface{}{fmt.Sprintf("[%s]", s.name)}, a...)...)
}
func (s *scheduler) WarnTag(tag string, a ...interface{}) {
	warn(append([]interface{}{fmt.Sprintf("[%s]", s.name), fmt.Sprintf("[%s]", tag)}, a...)...)
}
func (s *scheduler) Error(a ...interface{}) {
	err(append([]interface{}{fmt.Sprintf("[%s]", s.name)}, a...)...)
}
func (s *scheduler) ErrorTag(tag string, a ...interface{}) {
	err(append([]interface{}{fmt.Sprintf("[%s]", s.name), fmt.Sprintf("[%s]", tag)}, a...)...)
}
