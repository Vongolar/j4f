package jauthority

import (
	"JFFun/data/Dcommand"
)

//Authority 用户权限类型
type Authority = int

const (
	//Temp 没有鉴权的临时用户
	Temp Authority = 1
	//Player 玩家
	Player Authority = 1 << 1
	//Guest 游客玩家
	Guest Authority = 1 << 2
	//Root 最高权限
	Root Authority = 1 << 32
)

//WithAuthority 是否有权限
func WithAuthority(cmd Dcommand.Command, auth Authority) bool {
	if p, ok := commands[cmd]; ok {
		return p&auth != 0
	}
	return false
}
