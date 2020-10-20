package task

import (
	Jcommand "JFFun/data/command"
	Jerror "JFFun/data/error"
)

//Task 模块之间传递信息载体
type Task struct {
	Command   Jcommand.Command //原始任务命令
	AccountID string           //用户id，空代表是临时用户
	Data      interface{}      //消息上下文
	Request   *ChannelRequest
}

//Reply 回应任务
func (task *Task) Reply(errCode Jerror.Error, data interface{}) error {
	return nil
}

//OK 任务成功回应
func (task *Task) OK(resp interface{}) error {
	return task.Reply(Jerror.Error_ok, resp)
}

//Error 任务失败回应
func (task *Task) Error(errCode Jerror.Error) error {
	return task.Reply(errCode, nil)
}

//ResponseData 任务回应数据
type ResponseData struct {
	ErrCode Jerror.Error
	Data    interface{}
}
