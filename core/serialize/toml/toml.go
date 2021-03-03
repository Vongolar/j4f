package toml

import (
	"io"

	"github.com/BurntSushi/toml"
)

func GetExt() []string {
	return []string{`toml`}
}

func Encode(w io.Writer, v interface{}) error {
	return toml.NewEncoder(w).Encode(v)
}

func Decode(r io.Reader, out interface{}) error {
	_, err := toml.DecodeReader(r, out)
	return err
}
