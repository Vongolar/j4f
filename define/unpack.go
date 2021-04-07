package define

import "j4f/data/command"

func IsInvaildCMD(cmd command.Command) bool {
	return cmd == command.Command_invaild
}

func GetRequestStructByCMD(cmd command.Command) interface{} {
	switch cmd {
	default:
		return nil
	}
}
