package server

import (
	Jlog "JFFun/log"
	Jtag "JFFun/log/tag"
)

//Run 服务器启动
func Run() {
	if err := parseFlag(); err != nil {
		Jlog.Error(Jtag.Server, "服务器配置错误", err)
		return
	}

	if err := parseConfig(); err != nil {
		Jlog.Error(Jtag.Server, "服务器配置错误", err)
		return
	}
}
