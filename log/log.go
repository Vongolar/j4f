package log

import (
	"fmt"
)

func Info(tag string, a ...interface{}) {
	fmt.Println(a...)
}

func Warning(tag string, a ...interface{}) {
	fmt.Println(a...)
}

func Error(tag string, a ...interface{}) {
	fmt.Println(a...)
}
