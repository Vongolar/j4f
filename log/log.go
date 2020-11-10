package jlog

import (
	"fmt"
	"time"

	"github.com/logrusorgru/aurora/v3"
)

//Info 日志
func Info(tag string, a ...interface{}) {
	fmt.Print(aurora.Green(`[INFO] `))
	if len(tag) > 0 {
		fmt.Print(aurora.Green(fmt.Sprintf("[%s] ", tag)))
	}
	fmt.Print(a...)
	fmt.Println(aurora.Gray(10, fmt.Sprintf("	%s", time.Now().Format(time.RFC3339Nano))))
}

//Warning 日志
func Warning(tag string, a ...interface{}) {
	fmt.Print(aurora.Yellow(`[WARNING] `))
	if len(tag) > 0 {
		fmt.Print(aurora.Yellow(fmt.Sprintf("[%s] ", tag)))
	}
	fmt.Print(a...)
	fmt.Println(aurora.Gray(10, fmt.Sprintf("	%s", time.Now().Format(time.RFC3339Nano))))
}

//Error 日志
func Error(tag string, a ...interface{}) {
	fmt.Print(aurora.Red(`[Error] `))
	if len(tag) > 0 {
		fmt.Print(aurora.Red(fmt.Sprintf("[%s] ", tag)))
	}
	fmt.Print(a...)
	fmt.Println(aurora.Gray(10, fmt.Sprintf("	%s", time.Now().Format(time.RFC3339Nano))))
}
