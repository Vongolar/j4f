package jgate

import (
	jauthority "JFFun/authority"
	"JFFun/data/Dcommand"
	"JFFun/data/Dcommon"
	"JFFun/data/Derror"
	jtable "JFFun/database/table"
	jschedule "JFFun/schedule"
	jtask "JFFun/task"
	"database/sql"
	"fmt"
	"time"
)

//GetHandler 路由
func (m *MGate) GetHandler() map[Dcommand.Command]func(*jtask.Task) {
	return map[Dcommand.Command]func(*jtask.Task){
		Dcommand.Command_ping:          m.pong,
		Dcommand.Command_newRequest:    m.onNewRequest,
		Dcommand.Command_getPlayerAuth: m.getPlayerAuth,
		Dcommand.Command_bindConn:      m.onBindConn,
		Dcommand.Command_unbindConn:    m.onUnbindConn,
		Dcommand.Command_sync:          m.syncToPlayer,
	}
}

func (m *MGate) pong(task *jtask.Task) {
	task.OK()
}

func (m *MGate) onNewRequest(task *jtask.Task) {
	request, ok := task.Data.(*Dcommon.Request)
	if !ok {
		task.Error(Derror.Error_badRequest)
		return
	}

	if len(task.PlayerID) == 0 {
		//没有登陆的临时请求
		if jauthority.WithAuthority(request.Cmd, jauthority.Temp) {
			jschedule.HandleTask(request.Cmd, &jtask.Task{
				Request: task.Request,
				Raw:     request.Data,
			})
			return
		}
		task.Error(Derror.Error_noAuthority)
		return
	}

	if acc, ok := m.accounts[task.PlayerID]; ok {
		//玩家在这个gate上
		acc.lastActiveTS = time.Now()
		if jauthority.WithAuthority(request.Cmd, acc.auth) {
			jschedule.HandleTask(request.Cmd, &jtask.Task{
				PlayerID: acc.id,
				Request:  task.Request,
				Raw:      request.Data,
			})
			return
		}
		task.Error(Derror.Error_noAuthority)
		return
	}

	//询问所有gate模块用户的权限
	creq := jtask.NewInnerRequest()
	askTask := &jtask.Task{
		Request: creq,
		Data:    task.PlayerID,
	}
	jschedule.HandleTaskByOthers(m.name, Dcommand.Command_getPlayerAuth, askTask)
	if err, cresp := creq.Wait(); err == Derror.Error_ok {
		mresp := cresp.(jschedule.MutliResponse)
		for _, v := range mresp {
			cauth := v.(jauthority.Authority)
			if jauthority.WithAuthority(request.Cmd, cauth) {
				jschedule.HandleTask(request.Cmd, &jtask.Task{
					PlayerID: task.PlayerID,
					Request:  task.Request,
					Raw:      request.Data,
				})
				return
			}
			task.Error(Derror.Error_noAuthority)
			return
		}
	}

	//玩家绑定到gate
	//数据库读取用户权限
	row := m.getDB().QueryRow(fmt.Sprintf("select `auth` from `%s` where `id` = ?", jtable.Account), task.PlayerID)
	var auth jauthority.Authority
	err := row.Scan(&auth)
	if err != nil && err != sql.ErrNoRows {
		task.Error(Derror.Error_server)
		return
	}
	if err == sql.ErrNoRows {
		task.Error(Derror.Error_noAccount)
		return
	}

	m.bind(task.PlayerID, auth)
	m.onNewRequest(task)
}

func (m *MGate) getPlayerAuth(task *jtask.Task) {
	playerID, is := task.Data.(string)
	if !is {
		task.Error(Derror.Error_badRequest)
		return
	}
	if acc, ok := m.accounts[playerID]; ok {
		acc.lastActiveTS = time.Now()
		task.Reply(Derror.Error_ok, acc.auth)
		return
	}
	task.Error(Derror.Error_noAccount)
}

func (m *MGate) onBindConn(task *jtask.Task) {
	conn := task.Data.(iconnect)
	if acc, ok := m.accounts[task.PlayerID]; ok {
		//玩家在这个gate上
		acc.bindConn(conn)
		task.OK()
		return
	}

	//询问所有gate模块用户的权限
	creq := jtask.NewInnerRequest()
	askTask := &jtask.Task{
		Request: creq,
		Data:    task.PlayerID,
	}
	jschedule.HandleTaskByOthers(m.name, Dcommand.Command_getPlayerAuth, askTask)
	if err, cresp := creq.Wait(); err == Derror.Error_ok {
		mresp := cresp.(jschedule.MutliResponse)
		for name := range mresp {
			//指定到绑定模块
			jschedule.HandleTaskBy(name, Dcommand.Command_bindConn, task)
			return
		}
	}

	//玩家绑定到gate
	//数据库读取用户权限
	row := m.getDB().QueryRow(fmt.Sprintf("select `auth` from `%s` where `id` = ?", jtable.Account), task.PlayerID)
	var auth jauthority.Authority
	err := row.Scan(&auth)
	if err != nil && err != sql.ErrNoRows {
		task.Error(Derror.Error_server)
		return
	}
	if err == sql.ErrNoRows {
		task.Error(Derror.Error_noAccount)
		return
	}

	m.bind(task.PlayerID, auth)
	m.accounts[task.PlayerID].bindConn(conn)
	task.OK()
}

func (m *MGate) onUnbindConn(task *jtask.Task) {
	if acc, ok := m.accounts[task.PlayerID]; ok {
		acc.unbindConn()
	}
	task.OK()
}

func (m *MGate) syncToPlayer(task *jtask.Task) {
	defer task.OK()
	syncdata, ok := task.Data.(*jschedule.SyncData)
	if !ok {
		return
	}

	acc, ok := m.accounts[task.PlayerID]
	if !ok {
		return
	}

	acc.sync(syncdata.Scmd, syncdata.Data)
}
