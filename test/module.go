package test

import "log"

type Module struct {
}

func (m *Module) OnStart() {
	log.Println("测试模块打开")
}

func (m *Module) OnClose() {
	log.Println("测试模块关闭")
}
