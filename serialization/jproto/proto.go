package jproto

import (
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

func Encode(w io.Writer, v interface{}) error {
	b, err := proto.Marshal(v.(proto.Message))
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func Deconde(r io.Reader, v interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return proto.Unmarshal(b, v.(proto.Message))
}
