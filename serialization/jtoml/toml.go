package jtoml

import (
	"io"

	"github.com/BurntSushi/toml"
)

func GetExt() string {
	return ".toml"
}

func Encode(w io.Writer, v interface{}) error {
	encoder := toml.NewEncoder(w)
	return encoder.Encode(v)
}

func Deconde(r io.Reader, v interface{}) error {
	_, err := toml.DecodeReader(r, v)
	return err
}
