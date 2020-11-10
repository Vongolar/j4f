package jgame

import (
	jgamecore "JFFun/games/core"
)

//Game 游戏
type Game interface {
	Ready(id string, players []string) *jgamecore.Core //游戏准备工作
	Start()                                            //开始运行游戏
}
