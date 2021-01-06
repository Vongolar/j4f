package jjson

import (
	"encoding/json"
	"io"
)

func GetExt() string {
	return ".json"
}

func Encode(w io.Writer, v interface{}) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(v)
}

func Deconde(r io.Reader, v interface{}) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(v)
}
