/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-22 16:04:11
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\config\config.go
 * @Date: 2021-01-14 10:22:11
 * @描述: 解析配置
 */
package config

import (
	"JFFun/serialization/jjson"
	"JFFun/serialization/jtoml"
	"JFFun/serialization/jyaml"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/**
 * @description:
 * @Author: xuxinwei
 * @param {string} file
 * @param {interface{}} out
 * @return {*}
 * @描述: 解析本地配置文件
 */
func ParseLocalConfig(file string, out interface{}) error {
	if filepath.IsAbs(file) {
		f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		return parseLocalConfig(f, out)
	}

	//猜测路径
	paths := []string{file, filepath.Join("./cfg", file), filepath.Join("./config", file)}
	for _, path := range paths {
		f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

		if os.IsNotExist(err) {
			continue
		}

		if err != nil {
			return err
		}

		return parseLocalConfig(f, out)
	}

	return fmt.Errorf(fmt.Sprintf("can't find config file %s.", file))
}

func parseLocalConfig(f *os.File, out interface{}) error {
	defer f.Close()
	return parseConfig(f, filepath.Ext(f.Name()), out)
}

func parseConfig(r io.Reader, ext string, out interface{}) error {
	switch ext {
	case jjson.GetExt():
		return jjson.Deconde(r, out)
	case jtoml.GetExt():
		return jtoml.Deconde(r, out)
	case jyaml.GetExt():
		return jyaml.Decode(r, out)
	default:
		return fmt.Errorf("not support to decode config with '%s' extension.", ext)
	}
}
