/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-19 14:45:11
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\toml\toml.go
 * @Date: 2021-02-19 14:40:07
 * @描述: 文件描述
 */

package toml

import (
	"io"

	"github.com/BurntSushi/toml"
)

func GetExt() []string {
	return []string{`toml`}
}

func Encode(w io.Writer, v interface{}) error {
	return toml.NewEncoder(w).Encode(v)
}

func Decode(r io.Reader, out interface{}) error {
	_, err := toml.DecodeReader(r, out)
	return err
}
