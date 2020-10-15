package gate

import (
	Jcommand "JFFun/data/command"
)

//长连接
type connect interface {
	wait()
	listen()
	sync(cmd Jcommand.Command, data []byte) error
	close() error
}
