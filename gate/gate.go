package gate

import (
	Jconfig "JFFun/config"
	Jerror "JFFun/data/error"
	"JFFun/serialize"
	"JFFun/task"
	"bytes"
	"context"
	"encoding/binary"
	"sync"
)

type M_Gate struct {
	cfg config

	accMgr       accountMgr
	accMgrLocker sync.Mutex

	taskChan chan *task.Task
}

func (m *M_Gate) GetName() string {
	return `gate`
}

func (m *M_Gate) Init() error {
	//解析配置
	if err := Jconfig.ParseModuleConfig(m.GetName(), &m.cfg); err != nil {
		return err
	}

	//初始化容量
	m.accMgr = accountMgr{
		pool: make(map[string]*account, m.cfg.FitPlayerCount),
	}
	m.taskChan = make(chan *task.Task, m.cfg.CommandBuffer)

	return nil
}

func (m *M_Gate) Run(ctx context.Context) {
	m.listen()
Listen:
	for {
		select {
		case <-ctx.Done():
			break Listen
		case task := <-m.taskChan:

		}
	}
}

//开启监听服务
func (m *M_Gate) listen() {
	m.listenCommand()
}

func (m *M_Gate) listenCommand() {
	if len(m.cfg.HTTP) > 0 {
		go listenHTTP(m.cfg.HTTP, m.onRequest)
	}
}

func (m *M_Gate) onRequest(authority string, task *task.Task, data []byte) {
	var err error
	task.Data, err = serialize.DecodeReq(task.CMD, task.SMode, data)
	if err != nil {
		task.Error(Jerror.Error_decodeError, nil)
	}
	m.taskChan <- task
}

// type command struct {
// 	id    int64
// 	cmd   Jcommand.Command
// 	smode serialize.SerializeMode
// 	acc   *account
// 	// response task.Response
// 	data []byte
// }

// func (m *M_Gate) listenConsole() {
// 	var cmd int
// 	var mode int
// 	var data string
// 	var id int64
// 	consoleFlag := flag.NewFlagSet("console", flag.ContinueOnError)
// 	consoleFlag.IntVar(&cmd, "c", -1, "command id")
// 	consoleFlag.IntVar(&mode, "s", int(serialize.JSON), "mode of serialize message")
// 	consoleFlag.StringVar(&data, "m", "", "command message")
// 	consoleFlag.Int64Var(&id, "i", 0, "id of command")
// 	listenConsole(func(msg string) {
// 		consoleFlag.Parse(strings.Split(msg, " "))

// 		if cmd < 0 {
// 			return
// 		}

// 		if !serialize.VaildMode(mode) {
// 			return
// 		}

// 		m.accMgrLocker.Lock()
// 		a := m.accMgr.getAccount(rootID)
// 		if a == nil {
// 			a = &account{
// 				id:   rootID,
// 				auth: authRoot,
// 			}
// 		}
// 		m.accMgr.pool[a.id] = a
// 		m.accMgrLocker.Unlock()
// 		m.commandChan <- &command{
// 			acc:      a,
// 			cmd:      Jcommand.Command(cmd),
// 			smode:    serialize.SerializeMode(mode),
// 			id:       id,
// 			data:     *(*[]byte)(unsafe.Pointer(&data)),
// 			response: new(consoleResp),
// 		}
// 	})
// }

// func (m *M_Gate) listenHTTP() {
// 	listenHTTP(m.cfg.HTTP, func(resp *httpResp) {
// 		cmdStr := resp.request.Header.Get("Command")
// 		protoStr := resp.request.Header.Get("Proto")
// 		idStr := resp.request.Header.Get("ID")
// 		authorization := resp.request.Header.Get("Authorization")
// 		id, err := strconv.ParseInt(idStr, 10, 64)
// 		if err != nil {
// 			resp.Reply(0, Jerror.Error_decodeError, nil)
// 			return
// 		}

// 		cmd, err := strconv.Atoi(cmdStr)
// 		if err != nil {
// 			resp.Reply(id, Jerror.Error_decodeError, nil)
// 			return
// 		}
// 		proto, err := strconv.Atoi(protoStr)
// 		if err != nil {
// 			resp.Reply(id, Jerror.Error_decodeError, nil)
// 			return
// 		}

// 		b := make([]byte, 1024)
// 		resp.request.Body.Read(b)
// 		m.onCommand(id, cmd, proto, authorization, b, resp)
// 	})
// }

// func (m *M_Gate) listenWebsocket() {
// 	listenWebsocket(m.cfg.Websocket, func(conn *websocket.Conn, r *http.Request) {
// 		if err := r.ParseForm(); err != nil {
// 			conn.Close()
// 			return
// 		}

// 		authorization := r.Form.Get("Authorization")
// 		if len(authorization) == 0 { //长连接一定要有认证
// 			conn.Close()
// 			return
// 		}

// 		m.accMgrLocker.Lock()
// 		acc := m.accMgr.getAccount(authorization)
// 		if acc == nil {
// 			acc = &account{
// 				id:   authorization,
// 				auth: authPlayer,
// 			}
// 			m.accMgr.addAccount(acc)
// 		}
// 		m.accMgrLocker.Unlock()
// 		go acc.listenWebsocket(conn)
// 	})
// }

// //用户接入(长连接)
// func (m *M_Gate) onAccecpt(authorization string, conn net.Conn) {

// }

// //单条消息(短连接)
// func (m *M_Gate) onCommand(id int64, cmd int, smode int, authorization string, data []byte, response task.Response) {
// 	if cmd < 0 {
// 		response.Reply(id, Jerror.Error_noHandler, nil)
// 		return
// 	}

// 	if !serialize.VaildMode(smode) {
// 		response.Reply(id, Jerror.Error_noSupportProto, nil)
// 	}

// 	m.accMgrLocker.Lock()
// 	var acc *account
// 	if len(authorization) == 0 {
// 		acc = m.accMgr.getTempAccount()
// 	} else {
// 		acc = m.accMgr.getAccount(authorization)
// 		if acc == nil {
// 			acc = &account{
// 				id:   authorization,
// 				auth: authPlayer, //赋予权限根据创建账号给
// 			}
// 			m.accMgr.addAccount(acc)
// 		}
// 	}
// 	m.accMgrLocker.Unlock()
// 	m.commandChan <- &command{
// 		acc:      acc,
// 		cmd:      Jcommand.Command(cmd),
// 		smode:    serialize.SerializeMode(smode),
// 		id:       id,
// 		data:     data,
// 		response: response,
// 	}
// }

func intToBytes(n int) []byte {
	data := int32(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func bytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int32
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}
