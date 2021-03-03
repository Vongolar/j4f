package yaml

import (
	"io"

	"gopkg.in/yaml.v3"
)

func GetExt() []string {
	return []string{`yaml`, `yml`}
}

func Encode(w io.Writer, v interface{}) error {
	return yaml.NewEncoder(w).Encode(v)
}

func Decode(r io.Reader, out interface{}) error {
	return yaml.NewDecoder(r).Decode(out)
}
