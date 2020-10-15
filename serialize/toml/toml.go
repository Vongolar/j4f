package toml

import (
	"github.com/BurntSushi/toml"
)

func DecodeFile(file string, v interface{}) error {
	_, err := toml.DecodeFile(file, v)
	return err
}
