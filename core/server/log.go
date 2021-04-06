package server

import (
	"fmt"
	"j4f/core/loglevel"
	"j4f/core/task"
	"j4f/data/command"
	dlog "j4f/data/log"
	"time"
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

var buffer []*dlog.LogMessage
var useBuffer bool = true

func log(level loglevel.Level, tag string, msg ...interface{}) {
	if level < defaultConfig.MinLogLevel {
		return
	}

	if len(msg) == 0 {
		return
	}

	li := &dlog.LogMessage{
		Level:              int32(level),
		Tag:                tag,
		Msg:                fmt.Sprint(msg...),
		TimeStampleNanoSec: time.Now().UnixNano(),
	}

	if defaultConfig.Log {
		fmt.Print(format(li))
	}

	if useBuffer {
		buffer = append(buffer, li)
		return
	}

	slog(li)
}

func closeLogBuffer() {
	if useBuffer {
		useBuffer = false

		for _, item := range buffer {
			slog(item)
		}

		buffer = nil
	}
}

func slog(li *dlog.LogMessage) {
	Handle(&task.Task{
		CMD:  command.Command_log,
		Data: li,
	})
}

func format(l *dlog.LogMessage) string {
	return fmt.Sprintln(l)
}
