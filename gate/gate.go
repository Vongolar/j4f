package gate

import (
	Jconfig "JFFun/config"
	Jcommand "JFFun/data/command"
	"JFFun/serialize"
	"JFFun/task"
	"context"
	"flag"
	"strings"
	"sync"
	"unsafe"
)

type M_Gate struct {
	cfg config

	accMgr       accountMgr
	accMgrLocker sync.Mutex

	commandChan chan *command
}

func (m *M_Gate) GetName() string {
	return `gate`
}

func (m *M_Gate) Init() error {
	if err := Jconfig.ParseModuleConfig(m.GetName(), &m.cfg); err != nil {
		return err
	}

	m.accMgr = accountMgr{
		pool: make(map[string]*account),
	}

	m.commandChan = make(chan *command, 10)

	return nil
}

func (m *M_Gate) GetHandlers() map[Jcommand.Command]func(task *task.Task) {
	return map[Jcommand.Command]func(task *task.Task){
		Jcommand.Command_getOnlinePlayerCount: m.getOnlinePlayerCount,
	}
}

func (m *M_Gate) Run(ctx context.Context) {
	if m.cfg.Console {
		go m.listenConsole()
	}

Listen:
	for {
		select {
		case <-ctx.Done():
			break Listen
		case c := <-m.commandChan:
			c.acc.onCommand(c)
		}
	}
}

type command struct {
	id      int64
	cmd     Jcommand.Command
	smode   serialize.SerializeMode
	acc     *account
	respone task.Response
	data    []byte
}

func (m *M_Gate) listenConsole() {
	var cmd int
	var mode int
	var data string
	var id int64
	consoleFlag := flag.NewFlagSet("console", flag.ContinueOnError)
	consoleFlag.IntVar(&cmd, "c", -1, "command id")
	consoleFlag.IntVar(&mode, "s", int(serialize.JSON), "mode of serialize message")
	consoleFlag.StringVar(&data, "m", "", "command message")
	consoleFlag.Int64Var(&id, "i", 0, "id of command")
	listenConsole(func(msg string) {
		consoleFlag.Parse(strings.Split(msg, " "))

		if cmd < 0 {
			return
		}

		if !serialize.Invaild(mode) {
			return
		}

		m.accMgrLocker.Lock()
		a := m.accMgr.getAccount(rootID)
		if a == nil {
			a = &account{
				id:   rootID,
				auth: authRoot,
			}
		}
		m.accMgr.pool[a.id] = a
		m.accMgrLocker.Unlock()
		m.commandChan <- &command{
			acc:     a,
			cmd:     Jcommand.Command(cmd),
			smode:   serialize.SerializeMode(mode),
			id:      id,
			data:    *(*[]byte)(unsafe.Pointer(&data)),
			respone: new(consoleResp),
		}
	})
}
