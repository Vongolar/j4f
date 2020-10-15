package gate

import (
	Jcommand "JFFun/data/command"
	Jlog "JFFun/log"
	"JFFun/serialize"
	"JFFun/server"
	"JFFun/task"
	"fmt"

	"github.com/gorilla/websocket"
)

const rootID = "root"

type account struct {
	id            string
	auth          authority
	websocketConn *websocket.Conn
}

//消息到达时
func (acc *account) onCommand(request *command) {
	//检查权限
	if !authorityVaid(request.cmd, acc.auth) {
		// request.response.Reply(request.id, Jerror.Error_permissionDenied, nil)
		return
	}

	//ping 心跳直接返回
	if request.cmd == Jcommand.Command_ping {
		// request.response.Reply(request.id, Jerror.Error_ok, nil)
		return
	}

	//解析消息内容
	req, err := serialize.DecodeReq(request.cmd, request.smode, request.data)
	if err != nil {
		Jlog.Error("", err)
		// request.response.Reply(request.id, Jerror.Error_decodeError, nil)
		return
	}

	//生成任务
	task := &task.Task{
		ID:    request.id,
		CMD:   request.cmd,
		SMode: request.smode,
		Data:  req,
	}

	Jlog.Info("", fmt.Sprintf("\nReq from %s\ncmd : %d\nid : %d\n", request.acc.id, request.cmd, request.id))
	//任务分发
	server.HandleTask(task.CMD, task)
}

func (acc *account) sync(command Jcommand.Command, data interface{}) {

}
