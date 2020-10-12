package gate

type accountMgr struct {
	pool map[string]*account
}

func (mgr *accountMgr) getAccount(playerID string) *account {
	acc, _ := mgr.pool[playerID]
	return acc
}

func (mgr *accountMgr) addAccount(acc *account) {
	mgr.pool[acc.id] = acc
}
