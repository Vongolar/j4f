/*
 * @Author: Vongola
 * @FilePath: /JFFun/scaffold/module/gen.go
 * @Date: 2020-12-26 15:22:38
 * @Description: Gen files of a new module.
 * @描述: 生成新模块文件。
 * @LastEditTime: 2020-12-26 16:14:57
 * @LastEditors: Vongola
 */
package module

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//generate

/**
 * @description:Gen files of a new module.
 * @描述: 生成新模块文件。
 * @param {string} name: New module name; 要生成的新模块名;
 * @param {string} path: Root path; 模块生成路径;
 */
func genModule(name string, path string) {
	err := os.Mkdir(path+name, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	ioutil.WriteFile(path+"/"+name+"/"+name+".go", []byte(getModulefile(name)), os.ModePerm)
	ioutil.WriteFile(path+"/"+name+"/handler.go", []byte(getHandlerfile(name)), os.ModePerm)
}

const moduleTemplate = `package %s
import "context"
type %s struct{}
func (m %s)Init(cfg string) error { return nil }
func (m %s)Run(ctx context.Context){}`

func getModulefile(name string) string {
	sname := "M_" + strings.ToUpper(name[:1]) + name[1:]
	return fmt.Sprintf(moduleTemplate, name, sname, sname, sname)
}

const moduleHandlerTemplate = `package %s`

func getHandlerfile(name string) string {
	return fmt.Sprintf(moduleHandlerTemplate, name)
}
