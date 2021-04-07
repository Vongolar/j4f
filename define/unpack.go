package define

import (
	"j4f/data/command"
	"j4f/data/common"
)

func IsInvaildCMD(cmd command.Command) bool {
	return cmd == command.Command_invaild
}

func GetRequestStructByCMD(cmd command.Command) interface{} {
	switch cmd {
	case command.Command_closeModule:
		return new(common.ModuleName)
	case command.Command_restartModule:
		return new(common.ModuleName)
	default:
		return nil
	}
}
