package gate

import (
	"sync"
)

type accountMgr struct {
	lock sync.Mutex
	pool sync.Map
}

func (mgr *accountMgr) getAccount(key string) *account {
	v, ok := mgr.pool.Load(key)
	if !ok {
		return nil
	}
	acc, ok := v.(*account)
	if !ok {
		return nil
	}
	return acc
}

func (mgr *accountMgr) getTempAccount() *account {
	return new(account)
}

func (mgr *accountMgr) addAccount(acc *account) {
	if mgr.getAccount(acc.getIdentification()) != nil {
		return
	}
	mgr.pool.Store(acc.getIdentification(), acc)
}
