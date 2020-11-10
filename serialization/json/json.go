package jjson

import (
	"encoding/json"
)

//Marshal 序列化json
func Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, nil
	}
	return json.Marshal(v)
}

//UnMarshal 反序列化json
func UnMarshal(b []byte, out interface{}) error {
	return json.Unmarshal(b, out)
}
