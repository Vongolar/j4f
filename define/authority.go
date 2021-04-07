package define

import "j4f/data/command"

type Auth int

const (
	Auth_Guest   = 1 << iota //未登录临时用户
	Auth_User                //登录用户
	Auth_Server              //其他服务器
	Auth_Console             //控制台，唯一，启动时传key
)

func HasAuthority(au Auth, cmd command.Command) bool {
	if cmd <= command.Command_innerMax {
		return false
	}

	a, ok := cmdAuthority[cmd]
	if !ok {
		return true
	}
	return a&au != 0
}

var cmdAuthority = map[command.Command]Auth{
	command.Command_closeModule:   Auth_Console,
	command.Command_restartModule: Auth_Console,
}
