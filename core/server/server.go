package server

import (
	"context"
	"flag"
	jconfig "j4f/core/config"
	"j4f/core/module"
	"j4f/core/schduler"
	"os"
	"os/signal"
	"sync"
)

var (
	scheduleModuler schduler.ISchedule
)

func Run(scheduleMod schduler.ISchedule, mods ...module.Module) {
	rootCfgPath := flag.String(`c`, `./config/root.toml`, `配置`)
	flag.Parse()
	err := jconfig.ParseFile(*rootCfgPath, &defaultConfig)
	if err != nil {
		Err("配置错误", err)
		return
	}

	if len(defaultConfig.ModuleConfigs) <= len(mods) {
		Err("模块配置文件数不足")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	err = scheduleMod.Init(ctx, defaultConfig.ModuleConfigs[0].Name, defaultConfig.ModuleConfigs[0].Config)
	if err != nil {
		Err("调度模块初始化错误", err)
		cancel()
		return
	}

	wg.Add(1)
	scheduleModuler = scheduleMod
	go func() {
		defer func() {
			if err := recover(); err != nil {
				Err(err)
			}
			wg.Done()
		}()
		scheduleMod.Run()
	}()

	if len(mods) > 0 {
		registModules(mods[1:]...)
	}

	closeSign := make(chan os.Signal)
	signal.Notify(closeSign, os.Interrupt, os.Kill)
	<-closeSign

	cancel()

	wg.Wait()
}

func registModules(mods ...module.Module) {
	for _, mod := range mods {
		scheduleModuler.RegistModule(mod)
	}

	scheduleModuler.RunModules()
}
