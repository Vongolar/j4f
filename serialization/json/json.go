package json

import (
	"encoding/json"
)

//Marshal json序列化
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

//Unmarshal json反序列化
func Unmarshal(raw []byte, v interface{}) error {
	return json.Unmarshal(raw, v)
}
