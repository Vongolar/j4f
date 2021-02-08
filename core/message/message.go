/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 16:48:14
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\message\message.go
 * @Date: 2021-02-08 15:36:30
 * @描述: 文件描述
 */

package message

import (
	"errors"
	"io"
	"j4f/core/json"
	"j4f/core/protobuf"
	"j4f/data"
	"j4f/mux"
)

type Type = int
type Seria = int

const (
	_ Seria = iota
	Portobuf
	Json
)

const DefaultSeria = Portobuf

var (
	ErrNoSupportFormat = errors.New("不支持序列化格式")
)

func Decode(s Seria, r io.Reader, cmd data.Command) (out interface{}, err error) {
	out = mux.GetDataByCommand(cmd)
	if out == nil {
		return
	}

	switch s {
	case Portobuf:
		err = protobuf.Decode(r, out)
	case Json:
		err = json.Decode(r, out)
	default:
		err = ErrNoSupportFormat
	}
	return
}

func Encode(s Seria, w io.Writer, data interface{}) error {
	switch s {
	case Portobuf:
		return protobuf.Encode(w, data)
	case Json:
		return json.Encode(w, data)
	}
	return ErrNoSupportFormat
}
