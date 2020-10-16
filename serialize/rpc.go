package serialize

import (
	"JFFun/data/command"
)

func getStruct(cmd command.Command) (interface{}, error) {
	switch cmd {
	case command.Command_ping:
		fallthrough
	case command.Command_getOnlinePlayerCount:
		return nil, nil
	}
	return nil, ErrNoCommandDecoder
}
