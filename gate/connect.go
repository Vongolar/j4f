package jgate

import (
	"JFFun/data/Dcommand"
	"JFFun/data/Dcommon"
	"JFFun/data/Derror"
	jserialization "JFFun/serialization"
	"errors"
)

type iconnect interface {
	sync(data []byte) error
	close() error
}

type msgType byte

const (
	msgRequest  msgType = 0 //请求
	msgResponse msgType = 1 //回应
	msgSync     msgType = 2 //通知
)

type connRequest struct {
	id   int32
	conn iconnect
}

func (req *connRequest) Reply(err Derror.Error, data interface{}) error {
	d, e := jserialization.Marshal(jserialization.DefaultMode, data)
	if e != nil {
		err = Derror.Error_server
	}

	b, e := jserialization.Marshal(jserialization.DefaultMode, &Dcommon.Response{
		Id:   req.id,
		Err:  err,
		Data: d,
	})
	if e != nil {
		err = Derror.Error_server
	}

	return req.conn.sync(append([]byte{byte(msgResponse)}, b...))
}

func (acc *account) sync(scmd Dcommand.SyncCommand, data interface{}) error {
	if acc.conn != nil {
		d, e := jserialization.Marshal(jserialization.DefaultMode, data)
		if e != nil {
			return e
		}

		b, e := jserialization.Marshal(jserialization.DefaultMode, &Dcommon.SyncData{
			Scmd: scmd,
			Data: d,
		})
		if e != nil {
			return e
		}

		return acc.conn.sync(append([]byte{byte(msgSync)}, b...))
	}
	return errors.New("account not bind connect")
}
