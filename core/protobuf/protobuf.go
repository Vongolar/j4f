/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 15:52:26
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\protobuf\protobuf.go
 * @Date: 2021-02-08 15:46:44
 * @描述: 文件描述
 */
package protobuf

import (
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

func Decode(r io.Reader, out interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return proto.Unmarshal(b, out.(proto.Message))
}

func Encode(w io.Writer, v interface{}) error {
	b, err := proto.Marshal(v.(proto.Message))
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
