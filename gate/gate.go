package gate

import (
	Jcommand "JFFun/data/command"
	Jtask "JFFun/task"
	"context"
)

//MGate gate模块
type MGate struct {
}

//Init 初始化模块
func (m *MGate) Init(cfg string) error {
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
func (m *MGate) Run(context.Context) {

}
