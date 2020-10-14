package gate

import (
	Jconfig "JFFun/config"
	Jcommand "JFFun/data/command"
	Jerror "JFFun/data/error"
	"JFFun/serialize"
	"JFFun/task"
	"context"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"unsafe"

	"github.com/gorilla/websocket"
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

func (m *M_Gate) Run(ctx context.Context) {
	if m.cfg.Console {
		go m.listenConsole()
	}

	if len(m.cfg.HTTP) > 0 {
		go m.listenHTTP()
	}

	if len(m.cfg.Websocket) > 0 {
		go m.listenWebsocket()
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
	id       int64
	cmd      Jcommand.Command
	smode    serialize.SerializeMode
	acc      *account
	response task.Response
	data     []byte
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

		if !serialize.VaildMode(mode) {
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
			acc:      a,
			cmd:      Jcommand.Command(cmd),
			smode:    serialize.SerializeMode(mode),
			id:       id,
			data:     *(*[]byte)(unsafe.Pointer(&data)),
			response: new(consoleResp),
		}
	})
}

func (m *M_Gate) listenHTTP() {
	listenHTTP(m.cfg.HTTP, func(resp *httpResp) {
		cmdStr := resp.request.Header.Get("Command")
		protoStr := resp.request.Header.Get("Proto")
		idStr := resp.request.Header.Get("ID")
		authorization := resp.request.Header.Get("Authorization")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			resp.Reply(0, Jerror.Error_decodeError, nil)
			return
		}

		cmd, err := strconv.Atoi(cmdStr)
		if err != nil {
			resp.Reply(id, Jerror.Error_decodeError, nil)
			return
		}
		proto, err := strconv.Atoi(protoStr)
		if err != nil {
			resp.Reply(id, Jerror.Error_decodeError, nil)
			return
		}

		b := make([]byte, 1024)
		resp.request.Body.Read(b)
		m.onCommand(id, cmd, proto, authorization, b, resp)
	})
}

func (m *M_Gate) listenWebsocket() {
	listenWebsocket(m.cfg.Websocket, func(conn *websocket.Conn, r *http.Request) {
		fmt.Println(r.RequestURI)
	})
}

//用户接入(长连接)
func (m *M_Gate) onAccecpt() {

}

//单条消息(短连接)
func (m *M_Gate) onCommand(id int64, cmd int, smode int, authorization string, data []byte, response task.Response) {
	if cmd < 0 {
		response.Reply(id, Jerror.Error_noHandler, nil)
		return
	}

	if !serialize.VaildMode(smode) {
		response.Reply(id, Jerror.Error_noSupportProto, nil)
	}

	m.accMgrLocker.Lock()
	var acc *account
	if len(authorization) == 0 {
		acc = m.accMgr.getTempAccount()
	} else {
		acc = m.accMgr.getAccount(authorization)
	}
	m.accMgrLocker.Unlock()
	m.commandChan <- &command{
		acc:      acc,
		cmd:      Jcommand.Command(cmd),
		smode:    serialize.SerializeMode(smode),
		id:       id,
		data:     data,
		response: response,
	}
}
