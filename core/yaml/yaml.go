/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-19 14:51:59
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\yaml\yaml.go
 * @Date: 2021-02-19 14:45:52
 * @描述: 文件描述
 */

package yaml

import (
	"io"

	"gopkg.in/yaml.v3"
)

func GetExt() []string {
	return []string{`yaml`, `yml`}
}

func Encode(w io.Writer, v interface{}) error {
	return yaml.NewEncoder(w).Encode(v)
}

func Decode(r io.Reader, out interface{}) error {
	return yaml.NewDecoder(r).Decode(out)
}
