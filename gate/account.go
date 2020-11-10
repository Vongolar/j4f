package jgate

import (
	jauthority "JFFun/authority"
	jlog "JFFun/log"
	jtag "JFFun/log/tag"
	"fmt"
	"time"
)

type account struct {
	auth         jauthority.Authority
	id           string
	lastActiveTS time.Time

	conn iconnect
}

func (acc *account) bindConn(conn iconnect) {
	acc.unbindConn()
	acc.conn = conn
	jlog.Info(jtag.Gate, fmt.Sprintf("用户 %s 开启连接", acc.id))
}

func (acc *account) unbindConn() {
	if acc.conn != nil {
		err := acc.conn.close()
		acc.conn = nil
		jlog.Info(jtag.Gate, fmt.Sprintf("用户 %s 断开连接", acc.id))
		if err != nil {
			jlog.Error(jtag.Gate, "断开连接", err)
		}
	}
}

func (acc *account) unbind() {
	acc.unbindConn()
}

func (m *MGate) bind(playerID string, auth jauthority.Authority) {
	if _, ok := m.accounts[playerID]; ok {
		return
	}
	m.accounts[playerID] = &account{
		auth:         auth,
		id:           playerID,
		lastActiveTS: time.Now(),
	}

	jlog.Info(jtag.Gate, fmt.Sprintf("用户 %s 绑定到 %s 模块", playerID, m.name))
}
