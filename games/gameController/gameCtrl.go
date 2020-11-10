package jgamecontroller

import (
	"JFFun/data/Dgame"
	jgame "JFFun/game"
	"JFFun/games/lostcities"
)

//CreateGame 创建游戏
func CreateGame(gametype Dgame.GameType) jgame.Game {
	switch gametype {
	case Dgame.GameType_LostCities:
		return new(lostcities.Game)
	default:
		return nil
	}
}
