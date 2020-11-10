package jprotobuf

import (
	"errors"

	"github.com/golang/protobuf/proto"
)

var errInterface = errors.New("Error Interface")

//Marshal 序列化protobuf
func Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, nil
	}
	if value, ok := v.(proto.Message); ok {
		return proto.Marshal(value)
	}
	return nil, errInterface
}

//UnMarshal 反序列化protobuf
func UnMarshal(b []byte, out interface{}) error {
	if value, ok := out.(proto.Message); ok {
		return proto.Unmarshal(b, value)
	}
	return errInterface
}
