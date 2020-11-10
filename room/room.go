package jroom

import (
	"JFFun/data/Dgame"
	jgame "JFFun/game"
	jgamecore "JFFun/games/core"
	jlog "JFFun/log"
	jtag "JFFun/log/tag"
	"context"
)

//MRoom 游戏房间模块
type MRoom struct {
	playerList map[string]*room
}

//Init 初始化
func (m *MRoom) Init(cfg string) error {
	m.playerList = make(map[string]*room)
	return nil
}

//Run 运行
func (m *MRoom) Run(ctx context.Context, name string) {

}

//游戏运行容器
type room struct {
	id       string
	gameType Dgame.GameType
	game     jgame.Game
	core     *jgamecore.Core
	mod      *MRoom
	players  []player
}

type player struct {
	ready bool
	id    string
}

func (r *room) goRun() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				jlog.Error(jtag.Game(r.id), "异常关闭", err)
			}

			for _, p := range r.players {
				delete(r.mod.playerList, p.id)
			}
		}()

		var players []string
		for _, p := range r.players {
			players = append(players, p.id)
		}
		r.core = r.game.Ready(r.id, players)
		jlog.Info(jtag.Game(r.id), "游戏开始")
		r.game.Start()
	}()
}
