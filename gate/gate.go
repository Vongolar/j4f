package gate

import (
	Jcommand "JFFun/data/command"
	Jerror "JFFun/data/error"
	Jrpc "JFFun/rpc"
	"JFFun/schedule"
	Jserialization "JFFun/serialization"
	Jtoml "JFFun/serialization/toml"
	Jtask "JFFun/task"
	"context"
)

//MGate gate模块
type MGate struct {
	cfg         config
	accountMgr  *accountManager
	requestChan chan *request
	acceptChan  chan *acceptRequest
	onConnClose chan *connCloseEvent
}

//Init 初始化模块
func (m *MGate) Init(cfg []byte) error {
	err := Jtoml.Unmarshal(cfg, &m.cfg)
	if err != nil {
		return err
	}

	m.accountMgr = &accountManager{
		pool: make(map[string]*account, m.cfg.FitAccount),
	}

	m.requestChan = make(chan *request, m.cfg.RequestBuffer)
	m.acceptChan = make(chan *acceptRequest, m.cfg.RequestBuffer)
	m.onConnClose = make(chan *connCloseEvent, 10)
	return nil
}

//GetName 模块名
func (m *MGate) GetName() string {
	return `gate`
}

//Run 运行
func (m *MGate) Run(ctx context.Context) {
	m.listenNet()

	for {
		select {
		case <-ctx.Done():
			return
		case request := <-m.requestChan:
			m.handleRequest(request)
		case connReq := <-m.acceptChan:
			m.accountMgr.onAccountAccept(connReq)
		case event := <-m.onConnClose:
			m.accountMgr.onAccountConnClose(event)
		}
	}
}

func (m *MGate) listenNet() {
	if len(m.cfg.HTTP) > 0 {
		go m.listenHTTP(m.cfg.HTTP)
	}
	if len(m.cfg.Websocket) > 0 {
		go m.listenWebsocket(m.cfg.Websocket)
	}
}

func (m *MGate) getAccountID(authorization string) (string, error) {
	//解析鉴权信息
	//...
	return authorization, nil
}

type acceptRequest struct {
	accountID  string
	conn       connect
	resultChan chan bool
}

type request struct {
	cmd   Jcommand.Command
	smode Jserialization.SerializateType
	Jtask.Request
	accountID string
	data      []byte
}

func (m *MGate) handleRequest(req *request) {
	acc := m.accountMgr.getAccount(req.accountID)

	if !authorityCommand(req.cmd, acc.auth) {
		req.Reply(&Jtask.ResponseData{ErrCode: Jerror.Error_permissionDenied})
		return
	}

	data, err := Jrpc.Decode(req.cmd, req.smode, req.data)
	if err != nil {
		req.Reply(&Jtask.ResponseData{ErrCode: Jerror.Error_decode})
		return
	}
	err = schedule.HandleTask(req.cmd, &Jtask.Task{
		AccountID: acc.id,
		Data:      data,
		Request:   req.Request,
	})
	if err != nil {
		req.Reply(&Jtask.ResponseData{ErrCode: Jerror.Error_commandNotAllow})
	}
}

type connCloseEvent struct {
	accountID string
	conn      connect
}
