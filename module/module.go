package module

import (
	"context"
	"j4f/command"
	"j4f/task"
	"reflect"
	"sync/atomic"
)

type TaskHandle func(*task.Task)

type IModule interface {
	GetHandlers() map[int]TaskHandle
	OnStart()
	OnClose()
}

const (
	STATE_IDLE    int64 = iota // 从未启动
	STATE_WARMUP               // 启动中
	STATE_RUNNING              // 正常运行
	STATE_CLOSING              // 正在关闭，不接受新消息
	STATE_CLOSED
)

type Module struct {
	mod     IModule
	handler chan *task.Task
	ctx     context.Context
	state   int64
}

func CreateModule(mod IModule, capacity int, ctx context.Context) *Module {
	return &Module{
		mod:     mod,
		handler: make(chan *task.Task, capacity),
		ctx:     ctx,
		state:   STATE_IDLE,
	}
}

func (m *Module) GetCommands() []int {
	handlers := m.mod.GetHandlers()
	commands := make([]int, 0, len(handlers))
	for cmd, _ := range handlers {
		commands = append(commands, cmd)
	}
	return commands
}

func (m *Module) setState(state int64) {
	atomic.StoreInt64(&m.state, state)
}

func (m *Module) GetState() int64 {
	return atomic.LoadInt64(&m.state)
}

func (m *Module) GetType() reflect.Type {
	return reflect.TypeOf(m.mod)
}

func (m *Module) Run() {
	defer func() {
		m.setState(STATE_CLOSED)
	}()

	m.setState(STATE_WARMUP)
	handlers := m.mod.GetHandlers()
	m.mod.OnStart()

	m.setState(STATE_RUNNING)
	for {
		select {
		case t := <-m.handler:
			if t == nil { // 管道关闭，不接受新消息，关闭服务器
				return
			}

			if handler, exist := handlers[t.CommandID]; exist {
				handler(t)
			} else {
				// TODO: 不存在对应处理的错误
			}
		case <-m.ctx.Done():
			m.close()
		}
	}
}

func (m *Module) Exec(t *task.Task) {
	state := m.GetState()
	if state >= STATE_CLOSING {
		// TODO: 正在关闭，回应错误
		t.Reply(nil, command.Err_MODULE_CLOSE)
		return
	}

	m.handler <- t
}

func (m *Module) close() {
	if m.GetState() >= STATE_CLOSING {
		return
	}
	m.setState(STATE_CLOSING)
	m.mod.OnClose()
	close(m.handler)
}
