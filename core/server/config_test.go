package server

import (
	"bytes"
	"fmt"
	"j4f/core/serialize/toml"
	"testing"
)

func Test_Gen(t *testing.T) {
	buff := new(bytes.Buffer)
	toml.Encode(buff, defaultConfig)
	fmt.Print(buff.String())
}
