package server

import (
	"JFFun/serialization/jtoml"
	"bytes"
	"fmt"
	"testing"
)

func Test_GenTomlConfig(t *testing.T) {
	cfg := &config{Modules: []moduleConfigFile{
		moduleConfigFile{Name: "gate", Path: "gate.go"},
		moduleConfigFile{Name: "gate1", Path: "gate1.go"},
		moduleConfigFile{Name: "gate2", Path: "gate2.go"},
	}}

	buff := new(bytes.Buffer)
	if err := jtoml.Encode(buff, cfg); err == nil {
		fmt.Println(buff.String())
	}
}
