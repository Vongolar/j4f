package gate

import (
	Jcommand "JFFun/data/command"
)

//长连接
type connect interface {
	sync(cmd Jcommand.Command, data []byte) error
	close() error
}
