package gate

import (
	Jconfig "JFFun/config"
	"JFFun/data/command"
	"context"
	"strconv"
	"strings"
	"sync"
)

type M_Gate struct {
	cfg config

	accMgr accountMgr
}

func (m *M_Gate) GetName() string {
	return `gate`
}

func (m *M_Gate) Init() error {
	if err := Jconfig.ParseModuleConfig(m.GetName(), &m.cfg); err != nil {
		return err
	}

	m.accMgr = accountMgr{
		pool: sync.Map{},
	}

	return nil
}

func (m *M_Gate) Run(context.Context) {
	if m.cfg.Console {
		go listenConsole(func(msg string) {
			msgs := strings.Split(msg, " ")

			cmd, err := strconv.Atoi(msgs[0])
			if err != nil {
				return
			}

			mode := serJSON
			if len(msgs) == 3 {
				ss, err := strconv.Atoi(msgs[1])
				if err == nil {
					mode = serializeMode(ss)
				} else {
					return
				}
			}

			m.accMgr.lock.Lock()
			a := m.accMgr.getAccount(accRoot)
			if a == nil {
				a = &account{
					mold: accRoot,
				}
				m.accMgr.addAccount(a)
			}
			m.accMgr.lock.Unlock()
			a.addMsg(command.Command(cmd), mode, []byte(msgs[len(msgs)-1]), new(consoleReq))
		})
	}
}
