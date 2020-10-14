package gate

import (
	Jcommand "JFFun/data/command"
)

type authority = uint8

const (
	authRoot authority = 1 << 7
	authTemp authority = 1
)

var cmdAuthority = map[Jcommand.Command]uint8{
	Jcommand.Command_ping: 0xff,
}

func authorityVaid(cmd Jcommand.Command, auth authority) bool {
	if a, ok := cmdAuthority[cmd]; ok {
		return a&auth != 0
	}
	return true
}
