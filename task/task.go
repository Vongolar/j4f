package task

import (
	"JFFun/data/command"
	Jerror "JFFun/data/error"
)

//Task 模块之间传递
type Task struct {
	CMD command.Command

	AccountID string //用户id，空代表是临时用户

	Data interface{} //消息上下文
	Request
}

//ResponseData 任务回应数据
type ResponseData struct {
	ErrCode Jerror.Error
	Data    interface{}
}

//Request 任务请求源
type Request interface {
	Reply(resp *ResponseData) error
}

//Reply 回应任务
func (task *Task) Reply(errCode Jerror.Error, data interface{}) error {
	return task.Request.Reply(&ResponseData{
		ErrCode: errCode,
		Data:    data,
	})
}

//OK 任务成功回应
func (task *Task) OK(resp interface{}) error {
	return task.Reply(Jerror.Error_ok, resp)
}

//Error 任务失败回应
func (task *Task) Error(errCode Jerror.Error) error {
	return task.Reply(errCode, nil)
}
