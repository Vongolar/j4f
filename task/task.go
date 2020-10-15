package task

import (
	"JFFun/data/command"
	Jerror "JFFun/data/error"
	"JFFun/serialize"
)

type Task struct {
	ID    int64
	CMD   command.Command
	SMode serialize.SerializeMode
	Data  interface{}
	Request
}

type Request interface {
	Reply(id int64, errCode Jerror.Error, data []byte) error
}

func (task *Task) OK(resp interface{}) error {
	return task.Error(Jerror.Error_ok, resp)
}

func (task *Task) Error(errCode Jerror.Error, resp interface{}) error {
	if resp == nil {
		return task.Reply(task.ID, errCode, nil)
	}
	data, err := serialize.EncodeResp(task.CMD, task.SMode, resp)
	if err != nil {
		return err
	}
	return task.Reply(task.ID, errCode, data)
}
