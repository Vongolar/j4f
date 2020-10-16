package serialize

import (
	"JFFun/data/command"
	"encoding/json"
	"errors"

	"github.com/golang/protobuf/proto"
)

//Mode 序列化类型
type Mode = int

const (
	_ Mode = iota
	//JSON json
	JSON
	//Protobuf google
	Protobuf
)

//ErrEncode 序列化错误
var ErrEncode = errors.New("Encode Error")

//ErrDecode 反序列化错误
var ErrDecode = errors.New("Deconde Error")

//ErrNoCommandDecoder 命令没有对应反序列化对象
var ErrNoCommandDecoder = errors.New("No Command Decoder")

//Encode 数据序列化
func Encode(mode Mode, data interface{}) ([]byte, error) {
	switch mode {
	case JSON:
		return json.Marshal(data)
	case Protobuf:
		if data == nil {
			return nil, nil
		}
		if pd, ok := data.(proto.Message); ok {
			return proto.Marshal(pd)
		}
	}
	return nil, ErrEncode
}

//Decode 命令数据反序列化
func Decode(cmd command.Command, mode Mode, raw []byte) (interface{}, error) {
	s, err := getStruct(cmd)
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, nil
	}

	switch mode {
	case JSON:
		err = json.Unmarshal(raw, s)
		return s, err
	case Protobuf:
		if pd, ok := s.(proto.Message); ok {
			err = proto.Unmarshal(raw, pd)
			return s, err
		}
	}
	return nil, ErrDecode
}
