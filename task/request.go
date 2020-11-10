package jtask

import (
	"JFFun/data/Derror"
	"errors"
	"sync"
)

//Request 请求
type Request interface {
	Reply(err Derror.Error, data interface{}) error
}

//ErrDupleReply 请求已经重复回应
var ErrDupleReply = errors.New("reques has been replied")

//NewInnerRequest 获取
func NewInnerRequest() *InnerRequest {
	return &InnerRequest{
		channel: make(chan *innerResponse),
	}
}

//InnerRequest 进程内
type InnerRequest struct {
	locker   sync.Mutex
	replied  bool
	waiting  bool
	channel  chan *innerResponse
	response *innerResponse
}

type innerResponse struct {
	err  Derror.Error
	data interface{}
}

//Reply 回应
func (req *InnerRequest) Reply(err Derror.Error, data interface{}) error {
	req.locker.Lock()
	defer req.locker.Unlock()

	if req.replied {
		return ErrDupleReply
	}

	req.replied = true

	req.response = &innerResponse{
		err:  err,
		data: data,
	}

	if req.waiting {
		req.channel <- req.response
	}

	return nil
}

//Wait 等待回应
func (req *InnerRequest) Wait() (Derror.Error, interface{}) {
	req.locker.Lock()

	if req.replied {
		defer req.locker.Unlock()
	} else {
		req.waiting = true
		req.locker.Unlock()
		<-req.channel
	}

	return req.response.err, req.response.data
}
