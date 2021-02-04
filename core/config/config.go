/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-04 17:17:00
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\config\config.go
 * @Date: 2021-02-04 16:48:14
 * @描述: 文件描述
 */
package config

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"j4f/core/toml"
)

var (
	ErrNoSupportFormat = errors.New("不支持序列化格式")
	ErrFindConfigFile  = errors.New("找不到配置文件")
)

func DecodeConfigFromFile(path string, out interface{}) error {
	if filepath.IsAbs(path) {
		f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		return decodeConfigFromFile(f, out)
	}

	//猜测路径
	paths := []string{path, filepath.Join("./cfg", path), filepath.Join("./config", path)}
	for _, path := range paths {
		f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

		if os.IsNotExist(err) {
			continue
		}

		if err != nil {
			return err
		}

		return decodeConfigFromFile(f, out)
	}

	return ErrFindConfigFile
}

func decodeConfigFromFile(f *os.File, out interface{}) error {
	defer f.Close()
	return Decode(filepath.Ext(f.Name()), f, out)
}

func Decode(ext string, r io.Reader, out interface{}) error {
	switch strings.TrimLeft(ext, ".") {
	case "toml":
		return toml.Decode(r, out)
	default:
		return ErrNoSupportFormat
	}
}

func Encode(ext string, w io.Writer, v interface{}) error {
	switch strings.TrimLeft(ext, ".") {
	case "toml":
		return toml.Encode(w, v)
	default:
		return ErrNoSupportFormat
	}
}
