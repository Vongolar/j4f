package server

import (
	Jlog "JFFun/log"
	Jtag "JFFun/log/tag"
	Jschedule "JFFun/schedule"
	"context"
	"os"
	"os/signal"
	"sync"
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

	if err := initModules(); err != nil {
		Jlog.Error(Jtag.Server, "模块初始化错误", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	go Jschedule.Run(ctx, &wg)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	cancel()

	wg.Wait()

	Jlog.Info(Jtag.Server, "服务器关闭")
}
