package jyaml

import (
	"io"

	"gopkg.in/yaml.v2"
)

func GetExt() string {
	return ".yaml"
}

func Encode(w io.Writer, v interface{}) error {
	encoder := yaml.NewEncoder(w)
	return encoder.Encode(v)
}

func Decode(r io.Reader, v interface{}) error {
	decoder := yaml.NewDecoder(r)
	return decoder.Decode(v)
}
