package server

import (
	"context"
	"fmt"
	"j4f/command"
	"j4f/module"
	"j4f/schedule"
	"j4f/task"
	"j4f/test"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Server struct {
}

func (s *Server) Run() {
	log.Println("sever start")
	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)

	s.startSchedule(ctx, wg)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan // 本地不安全关闭方式

	log.Println("sever closing")

	cancel()
	wg.Wait()

	log.Println("sever closed")
}

func (s *Server) startSchedule(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	// TODO: schedule 模块容量
	cap := 1000
	mod := module.CreateModule(new(schedule.Module), cap, ctx)
	go func() {
		defer wg.Done()
		mod.Run()
	}()

	schedule.SetScheduler(mod)
	schedule.SetExecFunc(mod.Exec)

	mods := []module.IModule{new(test.Module)}

	for _, m := range mods {
		t := &task.Task{
			CommandID: command.CMD_REGIST_MODULE,
			Data: &schedule.RegistModuleData{
				Module:   m,
				Capacity: 100, // 具体容量
				Ctx:      ctx,
			},
		}
		schedule.Exec(t)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("午时已到")
	schedule.Exec(&task.Task{
		CommandID: command.CMD_TEST,
	})
}
