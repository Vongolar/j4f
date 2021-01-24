/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-23 23:48:09
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: /JFFun/config/config.go
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
		return parseFileConfig(f, out)
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

		return parseFileConfig(f, out)
	}

	return fmt.Errorf(fmt.Sprintf("找不到配置文件 %s 。", file))
}

func parseFileConfig(f *os.File, out interface{}) error {
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
		return fmt.Errorf("不支持 %s 后缀的配置文件。", ext)
	}
}
