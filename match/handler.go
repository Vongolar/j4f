package jmatch

import (
	"JFFun/data/Dcommand"
	"JFFun/data/Derror"
	"JFFun/data/Dgame"
	jserialization "JFFun/serialization"
	jtask "JFFun/task"
)

//GetHandler 路由
func (m *MMatch) GetHandler() map[Dcommand.Command]func(*jtask.Task) {
	return map[Dcommand.Command]func(*jtask.Task){
		Dcommand.Command_match:         m.match,
		Dcommand.Command_matchCancel:   m.cancelMatch,
		Dcommand.Command_getMatchState: m.getState,
		Dcommand.Command_matchSure:     m.sure,
	}
}

//单人匹配
func (m *MMatch) match(task *jtask.Task) {
	req := new(Dgame.MatchReq)
	if err := jserialization.UnMarshal(jserialization.DefaultMode, task.Raw, req); err != nil {
		task.Error(Derror.Error_badRequest)
		return
	}

	if _, ok := m.pool[req.Game]; !ok {
		task.Error(Derror.Error_badRequest)
		return
	}

	m.locker.Lock()
	if _, ok := m.players[task.PlayerID]; ok {
		task.Error(Derror.Error_inMatching)
		m.locker.Unlock()
		return
	}
	m.locker.Unlock()

	task.OK()
	m.channel <- act{
		cmd:      Dcommand.Command_match,
		playerID: task.PlayerID,
		data:     req,
	}
}

//cancelMatch 取消匹配
func (m *MMatch) cancelMatch(task *jtask.Task) {
	m.locker.Lock()

	if _, ok := m.players[task.PlayerID]; !ok {
		task.Error(Derror.Error_badRequest)
		m.locker.Unlock()
		return
	}
	m.locker.Unlock()
	task.OK()
	m.channel <- act{
		cmd:      Dcommand.Command_matchCancel,
		playerID: task.PlayerID,
	}
}

//getState 获取匹配状态
func (m *MMatch) getState(task *jtask.Task) {
	m.locker.Lock()
	defer m.locker.Unlock()

	if playerState, ok := m.players[task.PlayerID]; !ok {
		task.Reply(Derror.Error_ok, &Dgame.MatchStateInfo{
			State: Dgame.MatchState_noMatch,
		})
	} else {
		task.Reply(Derror.Error_ok, &Dgame.MatchStateInfo{
			State:          playerState.state,
			Game:           playerState.gt,
			StartTS:        playerState.startTS.Unix(),
			SureDeadLineTS: playerState.sureDeadline.Unix(),
		})
	}
}

//sure 确认匹配
func (m *MMatch) sure(task *jtask.Task) {
	m.locker.Lock()
	if state, ok := m.players[task.PlayerID]; !ok || state.state != Dgame.MatchState_waitSure {
		task.Error(Derror.Error_badRequest)
		m.locker.Unlock()
		return
	}
	m.locker.Unlock()
	task.OK()
	m.channel <- act{
		cmd:      Dcommand.Command_matchSure,
		playerID: task.PlayerID,
	}
}
