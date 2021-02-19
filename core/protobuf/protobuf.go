/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-19 15:05:15
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\protobuf\protobuf.go
 * @Date: 2021-02-19 14:55:19
 * @描述: 文件描述
 */
package protobuf

import (
	"errors"

	"github.com/golang/protobuf/proto"
)

var ErrNoProtoMessageType = errors.New("no proto message type.")

func Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return []byte{}, nil
	}

	if in, ok := v.(proto.Message); ok {
		return proto.Marshal(in)
	}

	return nil, ErrNoProtoMessageType
}

func Unmarshal(msg []byte, out interface{}) error {
	if len(msg) == 0 {
		out = nil
		return nil
	}

	if v, ok := out.(proto.Message); ok {
		return proto.Unmarshal(msg, v)
	}

	return ErrNoProtoMessageType
}
