package gate

import (
	Jcommand "JFFun/data/command"
)

type authorization uint32

const (
	temp   authorization = 1      //临时权限
	player authorization = 1 << 1 //玩家权限
)

var cmdAuthorization = map[Jcommand.Command]uint32{
	Jcommand.Command_ping: 0xffffffff,
}

func authorityCommand(cmd Jcommand.Command, a authorization) bool {
	if auth, ok := cmdAuthorization[cmd]; ok {
		return auth&uint32(a) != 0
	}
	return false
}
