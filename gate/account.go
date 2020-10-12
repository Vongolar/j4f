package gate

import (
	"JFFun/data/command"
	Jerror "JFFun/data/error"
	"JFFun/message"
)

type accountMold = string

const (
	accRoot accountMold = `root`
)

type account struct {
	mold accountMold
	id   string
}

func (acc *account) getIdentification() string {
	if len(acc.id) == 0 {
		return acc.mold
	}
	return acc.mold + "-" + acc.id
}

func (acc *account) addMsg(cmd command.Command, serMode serializeMode, msg []byte, request message.Request) {
	switch cmd {
	case command.Command_ping:
		request.Reply(Jerror.Error_ok, nil)
	}
}
