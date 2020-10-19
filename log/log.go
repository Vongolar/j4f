package log

import (
	"fmt"
)

//Info 消息
func Info(tag string, a ...interface{}) {
	fmt.Println(append([]interface{}{`[Info]`, tag}, a...)...)
}

//Warning 警告
func Warning(tag string, a ...interface{}) {
	fmt.Println(append([]interface{}{`[Warning]`, tag}, a...)...)
}

//Error 错误
func Error(tag string, a ...interface{}) {
	fmt.Println(append([]interface{}{`[Error]`, tag}, a...)...)
}
