package config

import (
	"JFFun/serialization/jjson"
	"JFFun/serialization/jtoml"
	"JFFun/serialization/jyaml"
	"fmt"
	"os"
	"path/filepath"
)

func ParseFileConfig(file string, out interface{}) error {
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	ext := filepath.Ext(file)

	switch ext {
	case jjson.GetExt():
		return jjson.Deconde(f, out)
	case jtoml.GetExt():
		return jtoml.Deconde(f, out)
	case jyaml.GetExt():
		return jyaml.Decode(f, out)
	}
	return fmt.Errorf("not support to decode file with extension '%s'", ext)
}
