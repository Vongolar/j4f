package jconfig

import (
	jtoml "JFFun/serialization/toml"
	"io/ioutil"
)

//LoadConfig 加载本地配置
func LoadConfig(path string, out interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return jtoml.UnMarshal(b, out)
}
