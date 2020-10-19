package toml

import (
	"github.com/BurntSushi/toml"
)

//Unmarshal 反序列化toml
func Unmarshal(raw []byte, v interface{}) error {
	return toml.Unmarshal(raw, v)
}
