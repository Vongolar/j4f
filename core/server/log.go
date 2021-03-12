package server

import (
	"j4f/core/loglevel"
	dlog "j4f/data/common/log"
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

var buffer []*dlog.Item
var useBuffer bool

func log(level loglevel.Level, tag string, msg ...interface{}) {
	if level < defaultConfig.MinLogLevel {
		return
	}

	if useBuffer {
		buffer = append(buffer, &dlog.Item{
			Level: int32(level),
			Tag:   tag,
			Msg:   msg,
			Ts:    time.Now().Unix(),
		})
		return
	}

	slog(level, tag, time.Now(), msg...)
}

func CloseLogBuffer() {
	if useBuffer {
		useBuffer = false

		for _, item := range buffer {
			slog(item.Level, item.Tag, item.TS, item.Msg...)
		}

		buffer = nil
	}
}

func slog(level loglevel.Level, tag string, time time.Time, msg ...interface{}) {

}
