/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 18:15:41
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
	"j4f/data"
)

type mod struct {
	M        module.Module
	C        chan *task.TaskHandleTuple
	Cfg      ModuleConfig
	handlers map[data.Command]task.Handler
}

type ModuleConfig struct {
	Name   string `toml:"name" yaml:"name" json:"name"`
	Config string `toml:"config" yaml:"config" json:"config"`
	Buffer int    `toml:"buffer" yaml:"buffer" json:"buffer"`
}
