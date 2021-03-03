package json

import (
	"encoding/json"
	"io"
)

func GetExt() []string {
	return []string{`json`}
}

func Encode(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func Decode(r io.Reader, out interface{}) error {
	return json.NewDecoder(r).Decode(out)
}
