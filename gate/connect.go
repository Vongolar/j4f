package gate

import (
	Jcommand "JFFun/data/command"
	Jserialization "JFFun/serialization"
	Jtask "JFFun/task"
	"encoding/binary"
	"errors"
)

type connect interface {
	sync(id uint32, cmd Jcommand.Command, resp *Jtask.ResponseData) error
	close() error
}

type connectRequest struct {
	id  uint32
	cmd Jcommand.Command
	connect
}

func (r *connectRequest) Reply(resp *Jtask.ResponseData) error {
	return r.sync(r.id, r.cmd, resp)
}

var errAnalysisBytesTooShort = errors.New("bytes is too short to analysis")

func analysisBytes(raw []byte) (id uint32, cmd Jcommand.Command, smode Jserialization.SerializateType, data []byte, err error) {
	if len(raw) < 8 {
		err = errAnalysisBytesTooShort
		return
	}
	id = binary.BigEndian.Uint32(raw[:4])
	cmd = Jcommand.Command(binary.BigEndian.Uint32(raw[4:8]))
	smode = Jserialization.SerializateType(raw[8])
	data = raw[9:]
	return
}
