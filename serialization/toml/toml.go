package jtoml

import (
	"unsafe"

	"github.com/BurntSushi/toml"
)

//UnMarshal 反序列化toml
func UnMarshal(b []byte, out interface{}) error {
	_, err := toml.Decode(*(*string)(unsafe.Pointer(&b)), out)
	return err
}
