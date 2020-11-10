package jserialization

import (
	jjson "JFFun/serialization/json"
	jprotobuf "JFFun/serialization/protobuf"
	"errors"
)

var DefaultMode = Protobuf

//Mode 序列化类型
type Mode = int

const (
	//Protobuf google protobuf
	Protobuf Mode = 0
	//JSON json
	JSON Mode = 1
)

var errNosupportMode = errors.New("no support serialization mode")

//Marshal 序列化
func Marshal(mode Mode, v interface{}) ([]byte, error) {
	if v == nil {
		return nil, nil
	}
	switch mode {
	case Protobuf:
		return jprotobuf.Marshal(v)
	case JSON:
		return jjson.Marshal(v)
	}
	return nil, errNosupportMode
}

//UnMarshal 反序列化
func UnMarshal(mode Mode, b []byte, out interface{}) error {
	switch mode {
	case Protobuf:
		return jprotobuf.UnMarshal(b, out)
	case JSON:
		return jjson.UnMarshal(b, out)
	}
	return errNosupportMode
}
