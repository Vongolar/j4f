/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 15:46:29
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\json\json.go
 * @Date: 2021-02-08 15:45:21
 * @描述: 文件描述
 */

package json

import (
	"encoding/json"
	"io"
)

func Decode(r io.Reader, out interface{}) error {
	return json.NewDecoder(r).Decode(out)
}

func Encode(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
