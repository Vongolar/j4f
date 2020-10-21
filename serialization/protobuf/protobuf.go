package protobuf

import (
	"errors"

	"github.com/golang/protobuf/proto"
)

//ErrNoProtoMessage 序列化/反序列化 对象不是proto.Message类型
var ErrNoProtoMessage = errors.New("type of v is not proto.Message")

//Marshal protobuf序列化
func Marshal(v interface{}) ([]byte, error) {
	if value, ok := v.(proto.Message); ok {
		return proto.Marshal(value)
	}
	return nil, ErrNoProtoMessage
}

//Unmarshal protobuf反序列化
func Unmarshal(raw []byte, v interface{}) error {
	if value, ok := v.(proto.Message); ok {
		return proto.Unmarshal(raw, value)
	}
	return ErrNoProtoMessage
}
