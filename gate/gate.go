package jgate

import (
	"JFFun/data/Dcommand"
	"JFFun/data/Derror"
	jlog "JFFun/log"
	jtag "JFFun/log/tag"
	jschedule "JFFun/schedule"
	jconfig "JFFun/serialization/config"
	jtask "JFFun/task"
	"context"
	"fmt"
	"time"
)

//MGate 网关服务器
type MGate struct {
	name     string
	cfg      config
	accounts map[string]*account
}

//Init 初始化
func (m *MGate) Init(cfg string) error {
	if err := jconfig.LoadConfig(cfg, &m.cfg); err != nil {
		return err
	}

	m.accounts = make(map[string]*account)

	return nil
}

//Run 运行
func (m *MGate) Run(ctx context.Context, name string) {
	m.name = name
	m.listen()
	ticker := time.NewTicker(time.Minute * time.Duration(m.cfg.ClearOfflineIntervalMin))
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			//定期清理不活跃
			m.clearOfflineAccounts()
		}
	}
}

func (m *MGate) listen() {
	if len(m.cfg.HTTP) > 0 {
		go func() {
			if err := m.listenHTTP(m.cfg.HTTP); err != nil {
				jlog.Error(jtag.Gate, "http listen error", err)
			}
		}()
	}

	if len(m.cfg.Websocket) > 0 {
		go func() {
			if err := m.listenWebsocket(m.cfg.Websocket); err != nil {
				jlog.Error(jtag.Gate, "websocket listen error", err)
			}
		}()
	}
}

func (m *MGate) clearOfflineAccounts() {
	now := time.Now()
	for _, acc := range m.accounts {
		if acc.lastActiveTS.Add(time.Minute * time.Duration(m.cfg.OfflineTimeoutMin)).Before(now) {
			m.unbind(acc)
		}
	}
}

func (m *MGate) unbind(acc *account) {
	delete(m.accounts, acc.id)
	acc.unbind()
	jlog.Info(jtag.Gate, fmt.Sprintf("%s 模块解绑用户 %s", m.name, acc.id))
}

func (m *MGate) authority(auth string) (Derror.Error, string) {
	req := jtask.NewInnerRequest()
	task := &jtask.Task{
		Request: req,
		Data:    auth,
	}
	jschedule.HandleTask(Dcommand.Command_authority, task)
	err, resp := req.Wait()
	if err != Derror.Error_ok {
		return err, ""
	}
	if id, ok := resp.(string); ok {
		return err, id
	}
	return Derror.Error_server, ""
}

func (m *MGate) bindConn(playerID string, conn iconnect) Derror.Error {
	req := jtask.NewInnerRequest()
	task := &jtask.Task{
		PlayerID: playerID,
		Data:     conn,
		Request:  req,
	}
	jschedule.HandleTaskBy(m.name, Dcommand.Command_bindConn, task)
	err, _ := req.Wait()
	return err
}

func (m *MGate) unbindConn(playerID string) {
	req := jtask.NewInnerRequest()
	task := &jtask.Task{
		PlayerID: playerID,
		Request:  req,
	}
	jschedule.HandleTaskByAll(Dcommand.Command_unbindConn, task)
	req.Wait()
}
