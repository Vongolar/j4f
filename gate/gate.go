package gate

import (
	Jcommand "JFFun/data/command"
	Jserialization "JFFun/serialization"
	Jtoml "JFFun/serialization/toml"
	Jtask "JFFun/task"
	"context"
)

//MGate gate模块
type MGate struct {
	cfg         config
	requestChan chan *request
}

//Init 初始化模块
func (m *MGate) Init(cfg []byte) error {
	err := Jtoml.Unmarshal(cfg, &m.cfg)
	if err != nil {
		return err
	}

	m.requestChan = make(chan *request, 10)

	return nil
}

//GetName 模块名
func (m *MGate) GetName() string {
	return `gate`
}

//GetHandlers 获取处理函数
func (m *MGate) GetHandlers() map[Jcommand.Command]func(task *Jtask.Task) {
	return nil
}

//Run 运行
func (m *MGate) Run(ctx context.Context) {
	m.listenNet()

	for {
		select {
		case <-ctx.Done():
			return
		case request := <-m.requestChan:
			request.Reply(&Jtask.ResponseData{})
		}
	}
}

func (m *MGate) listenNet() {
	if len(m.cfg.HTTP) > 0 {
		go m.listenHTTP(m.cfg.HTTP)
	}
}

func (m *MGate) getAccountID(authorization string) (string, error) {
	return authorization, nil
}

type request struct {
	cmd   Jcommand.Command
	smode Jserialization.SerializateType
	*Jtask.ChannelRequest
	accountID string
	data      []byte
}
