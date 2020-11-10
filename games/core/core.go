package jgamecore

import (
	"JFFun/data/Dcommand"
	"JFFun/data/Derror"
	"JFFun/data/Dgame"
	jschedule "JFFun/schedule"
	jserialization "JFFun/serialization"
	jtask "JFFun/task"
	"time"
)

//Core 游戏核心
type Core struct {
	getInfoChannel  chan *jtask.Task
	actChannel      chan *jtask.Task
	stateEndChannel chan *jtask.Task

	getGameInfoHandler func(*jtask.Task)
}

//GetInfoChannel 获取游戏信息channel
func (core *Core) GetInfoChannel() chan<- *jtask.Task {
	return core.getInfoChannel
}

//ActChannel 玩家行动channel
func (core *Core) ActChannel() chan<- *jtask.Task {
	return core.actChannel
}

//StateEndChannel 玩家状态结束channel
func (core *Core) StateEndChannel() chan<- *jtask.Task {
	return core.stateEndChannel
}

//Init 初始化
func (core *Core) Init(onGetGameInfo func(*jtask.Task)) {
	core.getGameInfoHandler = onGetGameInfo

	core.getInfoChannel = make(chan *jtask.Task)
	core.actChannel = make(chan *jtask.Task)
	core.stateEndChannel = make(chan *jtask.Task)
}

//WaitAct 等待玩家操作
func (core *Core) WaitAct(dur time.Duration, timeoutHandler func(), handler func(task *jtask.Task) bool, playerID ...string) {
	core.wait(dur, core.actChannel, timeoutHandler, func(task *jtask.Task) bool {
		if _, ok := task.Data.(*Dgame.Action); ok {
			return handler(task)
		}
		task.Error(Derror.Error_badRequest)
		return false
	}, playerID...)
}

//WaitStateEnd 等待玩家状态结束
func (core *Core) WaitStateEnd(dur time.Duration, state int32, playerID ...string) {
	core.wait(dur, core.stateEndChannel, nil, func(task *jtask.Task) bool {
		if v, ok := task.Data.(*Dgame.StateEnd); ok {
			task.OK()
			return v.State == state
		}
		task.Error(Derror.Error_badRequest)
		return false
	}, playerID...)
}

//SyncGameState 通知游戏状态
func (core *Core) SyncGameState(state int32, data interface{}, player ...string) {
	b, err := jserialization.Marshal(jserialization.DefaultMode, data)
	if err != nil {
		return
	}
	sd := &Dgame.SyncGameState{
		State: state,
		Data:  b,
	}
	for _, id := range player {
		jschedule.Sync(Dcommand.SyncCommand_gameState, sd, id)
	}
}

//SyncPlayerAct 通知玩家动作
func (core *Core) SyncPlayerAct(act int32, actPlayer string, data interface{}, player ...string) {
	b, err := jserialization.Marshal(jserialization.DefaultMode, data)
	if err != nil {
		return
	}
	sd := &Dgame.SyncGamePlayerAct{
		Act:    act,
		Data:   b,
		Player: actPlayer,
	}
	for _, id := range player {
		jschedule.Sync(Dcommand.SyncCommand_gamePlayerAct, sd, id)
	}
}

func (core *Core) wait(dur time.Duration, channel chan *jtask.Task, timeoutHandler func(), handler func(task *jtask.Task) bool, playerID ...string) {
	if len(playerID) == 0 {
		return
	}
	for {
		select {
		case <-time.After(dur * time.Second):
			if timeoutHandler != nil {
				timeoutHandler()
			}
			return

		case task := <-channel:
			index, player := -1, ""
			for index, player = range playerID {
				if player == task.PlayerID {
					break
				}
			}

			if index < 0 {
				task.Error(Derror.Error_noturn)
				continue
			}

			if handler(task) {
				playerID = append(playerID[:index], playerID[index+1:]...)
			}

			if len(playerID) == 0 {
				return
			}

		case task := <-core.getInfoChannel:
			if core.getGameInfoHandler != nil {
				core.getGameInfoHandler(task)
			}
		}
	}
}
