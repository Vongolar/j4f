package jauthority

import (
	"JFFun/data/Dcommand"
)

const (
	authtemp   = 0x000000001
	authplayer = 0x000000006
)

var commands = map[Dcommand.Command]int{
	Dcommand.Command_ping:          authplayer,
	Dcommand.Command_getGameInfo:   authplayer,
	Dcommand.Command_gameAct:       authplayer,
	Dcommand.Command_gameStateEnd:  authplayer,
	Dcommand.Command_guestLogin:    authtemp,
	Dcommand.Command_match:         authplayer,
	Dcommand.Command_matchCancel:   authplayer,
	Dcommand.Command_getMatchState: authplayer,
	Dcommand.Command_matchSure:     authplayer,
	Dcommand.Command_ready:         authplayer,
}
