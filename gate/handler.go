package gate

import (
	Jcommand "JFFun/data/command"
	"JFFun/data/common"
	"JFFun/task"
)

func (m *M_Gate) GetHandlers() map[Jcommand.Command]func(task *task.Task) {
	return map[Jcommand.Command]func(task *task.Task){
		Jcommand.Command_getOnlinePlayerCount: m.getOnlinePlayerCount,
	}
}

func (m *M_Gate) getOnlinePlayerCount(task *task.Task) {
	task.OK(&common.SingleInt64{
		Count: int64(len(m.accMgr.pool)),
	})
}
