package jroom

import (
	"JFFun/data/Dcommand"
	"JFFun/data/Derror"
	"JFFun/data/Dgame"
	"JFFun/data/Dinternal"
	jgamecontroller "JFFun/games/gameController"
	jlog "JFFun/log"
	jtag "JFFun/log/tag"
	jserialization "JFFun/serialization"
	jtask "JFFun/task"
	"time"
)

//GetHandler 路由
func (m *MRoom) GetHandler() map[Dcommand.Command]func(*jtask.Task) {
	return map[Dcommand.Command]func(*jtask.Task){
		Dcommand.Command_createGame:   m.createGame,
		Dcommand.Command_ready:        m.ready,
		Dcommand.Command_getGameInfo:  m.getGameInfo,
		Dcommand.Command_gameAct:      m.gameAct,
		Dcommand.Command_gameStateEnd: m.endGameState,
	}
}

func (m *MRoom) createGame(task *jtask.Task) {
	req, ok := task.Data.(*Dinternal.CreateGameContext)
	if !ok {
		task.Error(Derror.Error_badRequest)
		return
	}

	game := jgamecontroller.CreateGame(req.GameType)
	id := time.Now().Format(time.RFC822)
	r := &room{
		id:       id,
		gameType: req.GameType,
		game:     game,
		mod:      m,
	}

	for _, id := range req.Players {
		m.playerList[id] = r
		r.players = append(r.players, player{
			id:    id,
			ready: false,
		})
	}

	task.OK()
}

func (m *MRoom) ready(task *jtask.Task) {
	r, ok := m.playerList[task.PlayerID]
	if !ok {
		task.Error(Derror.Error_badRequest)
		return
	}

	task.OK()

	left := len(r.players)
	for i, player := range r.players {
		if player.id == task.PlayerID && !player.ready {
			r.players[i].ready = true
			jlog.Info(jtag.Match, player.id, "准备完毕")
			left--
		} else if player.id != task.PlayerID && player.ready {
			left--
		}
	}

	if left > 0 {
		return
	}

	jlog.Info(jtag.Match, "所有玩家准备完毕")
	r.goRun()
}

func (m *MRoom) getGameInfo(task *jtask.Task) {
	r, ok := m.playerList[task.PlayerID]
	if !ok {
		task.Error(Derror.Error_badRequest)
		return
	}
	r.core.GetInfoChannel() <- task
}

func (m *MRoom) gameAct(task *jtask.Task) {
	r, ok := m.playerList[task.PlayerID]
	if !ok {
		task.Error(Derror.Error_badRequest)
		return
	}

	req := new(Dgame.Action)
	if err := jserialization.UnMarshal(jserialization.DefaultMode, task.Raw, req); err != nil {
		task.Error(Derror.Error_badRequest)
		return
	}
	task.Data = req

	r.core.ActChannel() <- task
}

func (m *MRoom) endGameState(task *jtask.Task) {
	r, ok := m.playerList[task.PlayerID]
	if !ok {
		task.Error(Derror.Error_badRequest)
		return
	}
	req := new(Dgame.StateEnd)
	if err := jserialization.UnMarshal(jserialization.DefaultMode, task.Raw, req); err != nil {
		task.Error(Derror.Error_badRequest)
		return
	}
	task.Data = req
	r.core.StateEndChannel() <- task
}
