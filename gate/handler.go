package gate

import (
	Jcommand "JFFun/data/command"
	Jtask "JFFun/task"
)

//GetHandlers 获取处理函数
func (m *MGate) GetHandlers() map[Jcommand.Command]func(task *Jtask.Task) {
	return map[Jcommand.Command]func(task *Jtask.Task){
		Jcommand.Command_ping: m.pong,
	}
}

func (m *MGate) pong(task *Jtask.Task) {
	task.OK(nil)
}
