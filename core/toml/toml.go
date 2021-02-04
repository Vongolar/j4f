/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-04 16:34:34
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\toml\toml.go
 * @Date: 2021-02-04 16:31:12
 * @描述: 文件描述
 */
package toml

import (
	"io"

	"github.com/BurntSushi/toml"
)

func Decode(r io.Reader, out interface{}) error {
	_, err := toml.DecodeReader(r, out)
	return err
}

func Encode(w io.Writer, v interface{}) error {
	return toml.NewEncoder(w).Encode(v)
}
