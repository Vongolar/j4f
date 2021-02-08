/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 16:19:54
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\task\request.go
 * @Date: 2021-02-08 15:07:44
 * @描述: 文件描述
 */

package task

import (
	"io"
	"j4f/core/message"
	"j4f/data"
)

type Request interface {
	GetWriter() io.Writer
	Reply(err data.Error)
}

func (t *Task) OK(msg interface{}) {
	t.Reply(data.Error_ok, msg)
}

func (t *Task) Err(errCode data.Error) {
	t.Reply(errCode, nil)
}

func (t *Task) Reply(errCode data.Error, msg interface{}) {
	if msg == nil {
		t.Request.Reply(errCode)
		return
	}

	err := message.Encode(t.Seria, t.Request.GetWriter(), msg)
	if err != nil {
		return
	}
	t.Request.Reply(errCode)
}
