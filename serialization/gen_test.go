/*
 * @Author: Vongola
 * @FilePath: /JFFun/serialization/gen_test.go
 * @Date: 2021-01-24 00:26:20
 * @Description: file content
 * @描述: 文件描述
 * @LastEditTime: 2021-01-24 21:30:41
 * @LastEditors: Vongola
 */
package serialization

import (
	"JFFun/serialization/jtoml"
	"bytes"
	"testing"
)

type serverConfig struct {
	Name    string            `toml:"name" json:"name" yaml:"name"`
	Modules map[string]string `toml:"module" json:"module" yaml:"module"`
}

func Test_Gen(t *testing.T) {
	d := &serverConfig{
		Name: "server1",
		Modules: map[string]string{
			"m1": "momo",
			"m2": "mimi",
		},
	}
	w := new(bytes.Buffer)
	jtoml.Encode(w, d)
	t.Error(w.String())
}
