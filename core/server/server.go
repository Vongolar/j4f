package server

import (
	"context"
	"crypto/sha1"
	"flag"
	jconfig "j4f/core/config"
	"j4f/core/module"
	"j4f/core/scheduler"
	"os"
	"os/signal"
	"sync"
)

var (
	scheduleModuler scheduler.ISchedule

	hasConsoleKey = false
	consoleKey    [20]byte
)

func Run(scheduleMod scheduler.ISchedule, mods ...module.Module) {
	rootCfgPath := flag.String(`c`, `./config/root.toml`, `配置`)
	ck := flag.String(`k`, ``, `控制台key`)
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

	if len(*ck) > 0 {
		consoleKey = sha1.Sum([]byte(*ck))
		hasConsoleKey = true
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
		scheduleMod.Run(nil)
	}()

	if err := registModules(mods[0:]...); err != nil {
		cancel()
		return
	}

	closeSign := make(chan os.Signal)
	signal.Notify(closeSign, os.Interrupt, os.Kill)
	<-closeSign

	cancel()

	wg.Wait()
}

func registModules(mods ...module.Module) error {
	for i, mod := range mods {
		err := scheduleModuler.RegistModule(scheduler.ModuleWithCfg{
			Mod:    mod,
			Name:   defaultConfig.ModuleConfigs[i+1].Name,
			Config: defaultConfig.ModuleConfigs[i+1].Config,
		})

		if err != nil {
			return err
		}
	}

	scheduleModuler.RunModules()

	closeLogBuffer()
	return nil
}

func EqualConsoleKey(key string) bool {
	if !hasConsoleKey {
		return false
	}
	n := sha1.Sum([]byte(key))
	for i := 0; i < len(n); i++ {
		if n[i] != consoleKey[i] {
			return false
		}
	}

	return true
}
