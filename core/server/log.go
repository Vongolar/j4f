package server

import (
	"fmt"
	"j4f/core/loglevel"
)

func Info(msg ...interface{}) {
	InfoTag(``, msg...)
}

func InfoTag(tag string, msg ...interface{}) {
	log(loglevel.INFO, tag, msg...)
}

func Warn(msg ...interface{}) {
	WarnTag(``, msg...)
}

func WarnTag(tag string, msg ...interface{}) {
	log(loglevel.WARNING, tag, msg...)
}

func Err(msg ...interface{}) {
	ErrTag(``, msg...)
}

func ErrTag(tag string, msg ...interface{}) {
	log(loglevel.ERROR, tag, msg...)
}

func log(level loglevel.Level, tag string, msg ...interface{}) {
	if level < defaultConfig.MinLogLevel {
		return
	}

	fmt.Println(append([]interface{}{loglevel.GetLevelTag(level), tag}, msg...)...)
}
