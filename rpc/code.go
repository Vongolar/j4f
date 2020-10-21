package rpc

import (
	Jcommand "JFFun/data/command"
	Jserialization "JFFun/serialization"
	Jjson "JFFun/serialization/json"
	Jprotobuf "JFFun/serialization/protobuf"
	"errors"
)

//ErrNoSupportSerializateType 不支持序列化格式
var ErrNoSupportSerializateType = errors.New("No support the type for seriazlization")

//Encode 序列化
func Encode(st Jserialization.SerializateType, v interface{}) ([]byte, error) {
	if v == nil {
		return nil, nil
	}
	switch st {
	case Jserialization.JSON:
		return Jjson.Marshal(st)
	case Jserialization.Protobuf:
		return Jprotobuf.Marshal(st)
	}
	return nil, ErrNoSupportSerializateType
}

//Decode 反序列化
func Decode(cmd Jcommand.Command, st Jserialization.SerializateType, raw []byte) (interface{}, error) {
	if len(raw) == 0 {
		return nil, nil
	}

	v := getStructByCommand(cmd)
	if v == nil {
		return nil, nil
	}

	var err error
	switch st {
	case Jserialization.JSON:
		err = Jjson.Unmarshal(raw, v)
	case Jserialization.Protobuf:
		err = Jprotobuf.Unmarshal(raw, v)
	default:
		err = ErrNoSupportSerializateType
	}
	if err == nil {
		return v, nil
	}
	return nil, err
}
