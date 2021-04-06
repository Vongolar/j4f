package define

import "j4f/data/command"

func GetRequestStructByCMD(cmd command.Command) interface{} {
	switch cmd {
	default:
		return nil
	}
}
