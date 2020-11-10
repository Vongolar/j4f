package jserver

import (
	jredis "JFFun/database/redis"
	jsql "JFFun/database/sql"
	jlog "JFFun/log"
	jtag "JFFun/log/tag"
	jmodule "JFFun/module"
	jschedule "JFFun/schedule"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

var register map[string]func() jmodule.Module

//Regist 注册
func Regist(name string, creator func() jmodule.Module) {
	if register == nil {
		register = make(map[string]func() jmodule.Module)
	}
	register[name] = creator
}

//Run 服务器启动
func Run() {
	err := loadCfg()
	if err != nil {
		jlog.Error(jtag.Server, `配置`, err)
		return
	}

	if err = connectDB(); err != nil {
		jlog.Error(jtag.Server, `数据库连接`, err)
		return
	}

	if err = connectCache(); err != nil {
		jlog.Error(jtag.Server, `缓存数据库连接`, err)
		return
	}

	err = initModules()
	if err != nil {
		jlog.Error(jtag.Server, `初始化模块`, err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	jschedule.Run(ctx, &wg)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	cancel()
	wg.Wait()

	jsql.Close()
	jlog.Info(jtag.Server, `停止运行`)
}

func connectDB() error {
	for k, v := range cfg.Mysql {
		if err := jsql.ConnectDB(k, v.User, v.Password, v.Addr); err != nil {
			return err
		}
		jlog.Info(jtag.DataBase, fmt.Sprintf("%s 数据库连接成功", k))
	}
	return nil
}

func connectCache() error {
	for k, v := range cfg.Redis {
		if err := jredis.Connect(k, v.Addr, v.Password, v.DB); err != nil {
			return err
		}
		jlog.Info(jtag.Cache, fmt.Sprintf("%s 缓存数据库连接成功", k))
	}
	return nil
}

func initModules() error {
	for k, v := range cfg.Modules {
		if creator, ok := register[k]; ok {
			mod := creator()
			for i, path := range v.Paths {
				jlog.Info(jtag.Server, fmt.Sprintf("%s模块初始化", k), path)
				err := mod.Init(configPath + path)
				if err != nil {
					return err
				}
				jschedule.RegistModule(fmt.Sprintf("%s-%d", k, i), mod, v.Buf)
			}
		} else {
			jlog.Warning(jtag.Server, fmt.Sprintf("配置表中%s模块没有注册", k))
		}
	}
	return nil
}
