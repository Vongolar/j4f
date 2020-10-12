package gate

import (
	Jcommand "JFFun/data/command"
	Jerror "JFFun/data/error"
	Jlog "JFFun/log"
	"JFFun/rpc"
	"JFFun/server"
	"JFFun/task"
	"fmt"
)

const rootID = "root"

type account struct {
	id   string
	auth authority
}

func (acc *account) onCommand(request *command) {
	if request.cmd == Jcommand.Command_ping {
		request.respone.Reply(request.id, Jerror.Error_ok, nil)
		return
	}

	req, err := rpc.DecodeReq(request.cmd, request.smode, request.data)
	if err != nil {
		Jlog.Error(Jlog.TAG_CommandReq, err)
	}

	task := &task.Task{
		ID:       request.id,
		CMD:      request.cmd,
		SMode:    request.smode,
		Data:     req,
		Response: request.respone,
	}

	Jlog.Info(Jlog.TAG_CommandReq, fmt.Sprintf("\nReq from %s\ncmd : %d\nid : %d\n", request.acc.id, request.cmd, request.id))
	server.HandleTask(task.CMD, task)
}
