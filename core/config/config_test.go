/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-19 16:13:42
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\config\config_test.go
 * @Date: 2021-02-19 15:34:00
 * @描述: 文件描述
 */
package config

import "fmt"

type A struct {
	Name string `toml:"name" yaml:"name" json:"name"`
}

func Example_Json() {
	a := new(A)
	ParseFile("a.json", a)
	fmt.Println(a.Name)
	//Output: this is json config file
}

func Example_Yaml() {
	a := new(A)
	ParseFile("a.yml", a)
	fmt.Println(a.Name)
	//Output: this is yaml config file
}

func Example_Toml() {
	a := new(A)
	// ParseFile("E:/github/JFFun/core/config/a.toml", a)
	ParseFile("a.toml", a)
	fmt.Println(a.Name)
	//Output: this is toml config file
}
