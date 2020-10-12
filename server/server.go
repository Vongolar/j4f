package server

import (
	"JFFun/config"
	"JFFun/data/command"
	Jerror "JFFun/data/error"
	"JFFun/log"
	"JFFun/module"
	"JFFun/task"
	"context"
	"os"
	"os/signal"
	"sync"
)

func Run(modules ...module.Module) error {
	config.ParseServerConfig()

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	workHandlers = make(map[command.Command][]*workModule, len(modules))
	for i := 0; i < len(modules); i++ {
		if err := modules[i].Init(); err != nil {
			cancel()
			return err
		}
		handlers := modules[i].GetHandlers()
		wm := &workModule{
			module:   modules[i],
			taskChan: make(chan *taskForWork, 10),
			handlers: handlers,
		}
		workers = append(workers, wm)
		for cmd := range handlers {
			workHandlers[cmd] = append(workHandlers[cmd], wm)
		}
	}

	for _, worker := range workers {
		wg.Add(1)
		worker.goWait(ctx, &wg)
	}

	for _, mod := range modules {
		wg.Add(1)
		goRunModule(ctx, &wg, mod)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	cancel()

	wg.Wait()
	return nil
}

func HandleTask(command command.Command, task *task.Task) {
	if m, ok := workHandlers[command]; ok {
		if len(m) == 0 {
			goto NoHandler
		}

		//选择性能好的
		//...

		m[0].taskChan <- &taskForWork{
			cmd:  command,
			task: task,
		}

		return
	}
NoHandler:
	if err := task.Error(Jerror.Error_noHandler, nil); err != nil {
		log.Error(log.TAG_Server, err)
	}
}

func goRunModule(ctx context.Context, wg *sync.WaitGroup, mod module.Module) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error(log.TAG_Server, err)
			}
			wg.Done()
		}()
		mod.Run(ctx)
	}()
}

var workers []*workModule
var workHandlers map[command.Command][]*workModule

type workModule struct {
	module   module.Module
	taskChan chan *taskForWork
	handlers map[command.Command]func(task *task.Task)
}

type taskForWork struct {
	cmd  command.Command
	task *task.Task
}

func (mod *workModule) goWait(ctx context.Context, wg *sync.WaitGroup) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error(log.TAG_Server, err)
			}
			wg.Done()
		}()
		mod.wait(ctx)
	}()
}

func (mod *workModule) wait(ctx context.Context) {
Wait:
	for {
		select {
		case <-ctx.Done():
			break Wait
		case task := <-mod.taskChan:
			if handler, ok := mod.handlers[task.cmd]; ok {
				handler(task.task)
			} else {
				if err := task.task.Error(Jerror.Error_noHandler, nil); err != nil {
					log.Error(log.TAG_Server, err)
				}
			}
		}
	}
}
