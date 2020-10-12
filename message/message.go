package message

import (
	"JFFun/data/command"
	Jerror "JFFun/data/error"
)

type Message struct {
	CMD command.Command
	MSG interface{}
}

type Request interface {
	Reply(errCode Jerror.Error, msg []byte) error
}
