package jmatch

import (
	"JFFun/data/Dcommand"
	"JFFun/data/Derror"
	"JFFun/data/Dgame"
	"JFFun/data/Dinternal"
	jlog "JFFun/log"
	jtag "JFFun/log/tag"
	jschedule "JFFun/schedule"
	jconfig "JFFun/serialization/config"
	jtask "JFFun/task"
	"context"
	"fmt"
	"sync"
	"time"
)

//MMatch 匹配模块，目前人数凑满就行
type MMatch struct {
	cfg      config
	pool     map[Dgame.GameType][]string //匹配池
	waitList [][]string                  //等待确认

	locker  sync.Mutex
	players map[string]player

	channel chan act
}

type act struct {
	cmd      Dcommand.Command
	playerID string
	data     interface{}
}

type player struct {
	gt           Dgame.GameType
	state        Dgame.MatchState
	startTS      time.Time
	sureDeadline time.Time
}

//Init 初始化
func (m *MMatch) Init(cfg string) error {
	if err := jconfig.LoadConfig(cfg, &m.cfg); err != nil {
		return err
	}

	m.pool = make(map[Dgame.GameType][]string, len(m.cfg.Game))
	m.players = make(map[string]player)
	m.channel = make(chan act, m.cfg.Buff)

	for k := range m.cfg.Game {
		if t, ok := Dgame.GameType_value[k]; ok {
			m.pool[Dgame.GameType(t)] = make([]string, 0)
		} else {
			return fmt.Errorf("game %s is invaild", k)
		}
	}

	return nil
}

//Run 开始运行
func (m *MMatch) Run(ctx context.Context, name string) {
	ticker := time.NewTicker(time.Second)
	longTicker := time.NewTicker(time.Minute * time.Duration(m.cfg.ClearMatchIntervalMin))
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			m.checkDeadline()
		case <-longTicker.C:
			m.clear()
		case action := <-m.channel:
			switch action.cmd {
			case Dcommand.Command_match:
				m.addPlayer(action.data.(*Dgame.MatchReq).Game, action.playerID)
			case Dcommand.Command_matchCancel:
				m.subPlayer(action.playerID)
			case Dcommand.Command_matchSure:
				m.sureMatch(action.playerID)
			}
		}
	}
}

func (m *MMatch) addPlayer(gameType Dgame.GameType, playerID string) {
	m.locker.Lock()
	defer m.locker.Unlock()

	if _, exist := m.players[playerID]; exist {
		return
	}

	m.players[playerID] = player{
		gt:      gameType,
		state:   Dgame.MatchState_matching,
		startTS: time.Now(),
	}

	m.pool[gameType] = append(m.pool[gameType], playerID)

	jlog.Info(jtag.Match, playerID, "加入等待队列")

	fit := m.cfg.Game[Dgame.GameType_name[int32(gameType)]].Player

	if len(m.pool[gameType]) < fit {
		return
	}

	//满足开局条件
	choosers := m.pool[gameType][:fit]
	m.pool[gameType] = m.pool[gameType][fit:]

	m.waitList = append(m.waitList, choosers)

	for _, chooser := range choosers {
		m.players[chooser] = player{
			gt:           gameType,
			state:        Dgame.MatchState_waitSure,
			startTS:      m.players[chooser].startTS,
			sureDeadline: time.Now().Add(time.Duration(m.cfg.SureDurSec) * time.Second),
		}
		jlog.Info(jtag.Match, chooser, "进入匹配等待确认阶段")
		jschedule.Sync(Dcommand.SyncCommand_matchCheck, &Dgame.MatchStateInfo{
			State:          Dgame.MatchState_waitSure,
			Game:           gameType,
			StartTS:        m.players[chooser].startTS.Unix(),
			SureDeadLineTS: m.players[chooser].sureDeadline.Unix(),
		}, chooser)
	}
}

func (m *MMatch) subPlayer(playerID string) {
	m.locker.Lock()
	defer m.locker.Unlock()
	state, exist := m.players[playerID]
	if !exist {
		return
	}

	if state.state == Dgame.MatchState_matching {
		for i, id := range m.pool[state.gt] {
			if id == playerID {
				m.pool[state.gt] = append(m.pool[state.gt][:i], m.pool[state.gt][i+1:]...)
				break
			}
		}
		startTS := m.players[playerID].startTS
		gt := m.players[playerID].gt
		delete(m.players, playerID)
		jschedule.Sync(Dcommand.SyncCommand_matchFail, &Dgame.MatchStateInfo{
			State:   Dgame.MatchState_matching,
			Game:    gt,
			StartTS: startTS.Unix(),
		}, playerID)
		return
	}

	//玩家已在等待确认队列中
	var waitList []string

FindWaitList:
	for i, list := range m.waitList {
		for _, id := range list {
			if id == playerID {
				waitList = list
				m.waitList = append(m.waitList[:i], m.waitList[:i+1]...)
				break FindWaitList
			}
		}
	}

	//队列中的玩家匹配失败
	for _, id := range waitList {
		startTS := m.players[id].startTS
		gt := m.players[id].gt
		delete(m.players, id)
		jschedule.Sync(Dcommand.SyncCommand_matchFail, &Dgame.MatchStateInfo{
			State:   Dgame.MatchState_waitSure,
			Game:    gt,
			StartTS: startTS.Unix(),
		}, playerID)
	}
}

func (m *MMatch) sureMatch(playerID string) {
	m.locker.Lock()
	defer m.locker.Unlock()
	state, exist := m.players[playerID]
	if !exist || state.state != Dgame.MatchState_waitSure {
		return
	}
	m.players[playerID] = player{
		state:        Dgame.MatchState_sure,
		gt:           state.gt,
		startTS:      state.startTS,
		sureDeadline: state.sureDeadline,
	}

	for index, list := range m.waitList {
		inList := false
		for _, id := range list {
			if id == playerID {
				inList = true
				break
			}
		}

		allCheck := true
		if inList {
			for _, id := range list {
				if m.players[id].state != Dgame.MatchState_sure {
					allCheck = false
					break
				}
			}
			if allCheck {
				m.waitList = append(m.waitList[:index], m.waitList[index+1:]...)
				m.matchSuccess(list)
			}
			break
		}
	}

}

func (m *MMatch) matchSuccess(playerIDs []string) {
	gt := m.players[playerIDs[0]].gt

	req := jtask.NewInnerRequest()
	task := &jtask.Task{
		Request: req,
		Data: &Dinternal.CreateGameContext{
			GameType: gt,
			Players:  playerIDs,
		},
	}

	jschedule.HandleTask(Dcommand.Command_createGame, task)
	err, _ := req.Wait()

	if err != Derror.Error_ok {
		//匹配失败
		for _, id := range playerIDs {
			jschedule.Sync(Dcommand.SyncCommand_matchFail, &Dgame.MatchStateInfo{
				State:          m.players[id].state,
				Game:           m.players[id].gt,
				StartTS:        m.players[id].startTS.Unix(),
				SureDeadLineTS: m.players[id].sureDeadline.Unix(),
			}, id)
			delete(m.players, id)
		}
		return
	}

	//成功
	for _, id := range playerIDs {
		jschedule.Sync(Dcommand.SyncCommand_matchSuccess, &Dgame.MatchSucInfo{
			Players: playerIDs,
			Info: &Dgame.MatchStateInfo{
				State:          m.players[id].state,
				Game:           gt,
				StartTS:        m.players[id].startTS.Unix(),
				SureDeadLineTS: m.players[id].sureDeadline.Unix(),
			},
		}, id)
		delete(m.players, id)
	}
}

func (m *MMatch) checkDeadline() {
	m.locker.Lock()
	defer m.locker.Unlock()

	for index, list := range m.waitList {
		if m.players[list[len(list)-1]].sureDeadline.Add(time.Second).After(time.Now()) {
			//未超时
			continue
		}
		//超时
		jlog.Info(jtag.Match, "队列超时")
		m.waitList = append(m.waitList[:index], m.waitList[index+1:]...)

		for _, id := range list {
			jschedule.Sync(Dcommand.SyncCommand_matchFail, &Dgame.MatchStateInfo{
				State:          m.players[id].state,
				Game:           m.players[id].gt,
				StartTS:        m.players[id].startTS.Unix(),
				SureDeadLineTS: m.players[id].sureDeadline.Unix(),
			}, id)
			jlog.Info(jtag.Match, id, "超时移出匹配队列")
			delete(m.players, id)
		}
	}
}

func (m *MMatch) clear() {
	m.locker.Lock()
	defer m.locker.Unlock()
	now := time.Now()
	for id, state := range m.players {
		if state.state == Dgame.MatchState_matching && state.startTS.Add(time.Duration(m.cfg.ClearMatchSec)*time.Second).Before(now) {
			for i, playerID := range m.pool[state.gt] {
				if playerID == id {
					m.pool[state.gt] = append(m.pool[state.gt][:i], m.pool[state.gt][i+1:]...)
					jschedule.Sync(Dcommand.SyncCommand_matchFail, &Dgame.MatchStateInfo{
						State:          m.players[id].state,
						Game:           m.players[id].gt,
						StartTS:        m.players[id].startTS.Unix(),
						SureDeadLineTS: m.players[id].sureDeadline.Unix(),
					}, id)
					jlog.Info(jtag.Match, id, "匹配不到玩家")
					delete(m.players, id)
					break
				}
			}
		}
	}
}
