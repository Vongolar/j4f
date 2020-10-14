package gate

type accountMgr struct {
	pool map[string]*account
}

func (mgr *accountMgr) getAccount(playerID string) *account {
	if mgr.pool == nil {
		return nil
	}
	acc, _ := mgr.pool[playerID]
	return acc
}

func (mgr *accountMgr) getTempAccount() *account {
	return &account{
		auth: authTemp,
	}
}

func (mgr *accountMgr) addAccount(acc *account) {
	if mgr.pool == nil {
		mgr.pool = make(map[string]*account)
	}
	mgr.pool[acc.id] = acc
}
