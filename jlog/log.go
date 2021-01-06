package jlog

import (
	"fmt"

	"github.com/logrusorgru/aurora/v3"
)

func Info(a ...interface{}) {
	tmp := make([]interface{}, len(a)+1)
	tmp[0] = aurora.White("[INFO]")
	for i, item := range a {
		tmp[i+1] = aurora.White(item)
	}
	fmt.Println(tmp...)
}

func InfoWithTag(tag string, a ...interface{}) {
	tmp := make([]interface{}, len(a)+2)
	tmp[0] = aurora.White("[INFO]")
	tmp[1] = aurora.White(fmt.Sprintf("[%s]", tag))

	for i, item := range a {
		tmp[i+2] = aurora.White(item)
	}
	fmt.Println(tmp...)
}

func Warning(a ...interface{}) {
	tmp := make([]interface{}, len(a)+1)
	tmp[0] = aurora.Yellow("[WARN]")
	for i, item := range a {
		tmp[i+1] = aurora.Yellow(item)
	}
	fmt.Println(tmp...)
}

func WarningWithTag(tag string, a ...interface{}) {
	tmp := make([]interface{}, len(a)+2)
	tmp[0] = aurora.Yellow("[WARN]")
	tmp[1] = aurora.Yellow(fmt.Sprintf("[%s]", tag))

	for i, item := range a {
		tmp[i+2] = aurora.Yellow(item)
	}
	fmt.Println(tmp...)
}

func Error(a ...interface{}) {
	tmp := make([]interface{}, len(a)+1)
	tmp[0] = aurora.Red("[ERRO]")
	for i, item := range a {
		tmp[i+1] = aurora.Red(item)
	}
	fmt.Println(tmp...)
}

func ErrorWithTag(tag string, a ...interface{}) {
	tmp := make([]interface{}, len(a)+2)
	tmp[0] = aurora.Red("[ERRO]")
	tmp[1] = aurora.Red(fmt.Sprintf("[%s]", tag))

	for i, item := range a {
		tmp[i+2] = aurora.Red(item)
	}
	fmt.Println(tmp...)
}
