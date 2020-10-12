package gate

import (
	"JFFun/data/common"
	"JFFun/task"
)

func (m *M_Gate) getOnlinePlayerCount(task *task.Task) {
	task.OK(&common.SingleInt64{
		Count: int64(len(m.accMgr.pool)),
	})
}
