/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-04 19:35:35
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\server\mod.go
 * @Date: 2021-02-04 19:18:23
 * @描述: 文件描述
 */

package server

import (
	"j4f/core/module"
	"j4f/core/task"
)

type mod struct {
	M   module.Module
	C   chan *task.TaskHandleTuple
	Cfg ModuleConfig
}

type ModuleConfig struct {
	Name   string `toml:"name" yaml:"name" json:"name"`
	Config string `toml:"config" yaml:"config" json:"config"`
	Buffer int    `toml:"buffer" yaml:"buffer" json:"buffer"`
}
