package gate

type accountManager struct {
	pool map[string]*account
}

func (mgr *accountManager) getAccount(id string) *account {
	if len(id) == 0 {
		return &account{
			auth: temp,
		}
	}
	if acc, ok := mgr.pool[id]; ok {
		return acc
	}

	acc := &account{
		id:   id,
		auth: player,
	}
	mgr.pool[id] = acc
	return acc
}

func (mgr *accountManager) onAccountAccept(req *acceptRequest) {
	acc := mgr.getAccount(req.accountID)
	if acc.conn != req.conn {
		if acc.conn != nil {
			acc.conn.close()
		}
	}
	acc.conn = req.conn
	req.resultChan <- true
}

func (mgr *accountManager) onAccountConnClose(event *connCloseEvent) {
	if acc, ok := mgr.pool[event.accountID]; ok && acc.conn != nil && acc.conn == event.conn {
		acc.conn.close()
	}
}
