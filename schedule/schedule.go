package schedule

import (
	Jcommand "JFFun/data/command"
	Jlog "JFFun/log"
	Jtag "JFFun/log/tag"
	Jmodule "JFFun/module"
	Jtask "JFFun/task"
	"context"
	"errors"
	"fmt"
	"sync"
)

var workers map[string][]*worker
var handlers map[Jcommand.Command][]*worker

//Regist 注册模块
func Regist(module Jmodule.Module) {
	if workers == nil {
		workers = make(map[string][]*worker)
	}

	if handlers == nil {
		handlers = make(map[Jcommand.Command][]*worker)
	}

	w := &worker{
		mod:         module,
		taskChannel: make(chan *task, 6),
		handlers:    module.GetHandlers(),
	}

	for cmd := range w.handlers {
		if _, ok := handlers[cmd]; !ok {
			handlers[cmd] = make([]*worker, 0)
		}
		handlers[cmd] = append(handlers[cmd], w)
	}

	if _, ok := workers[module.GetName()]; !ok {
		workers[module.GetName()] = make([]*worker, 0)
	}

	workers[module.GetName()] = append(workers[module.GetName()], w)
}

//Run 运行
func Run(ctx context.Context, wg *sync.WaitGroup) {
	for name, works := range workers {
		Jlog.Info(Jtag.Schedule, fmt.Sprintf("注册 %s 模块 %d 个", name, len(works)))

		for _, work := range works {
			wg.Add(1)
			goRunModule(ctx, work.mod, wg)
		}

		for _, work := range works {
			wg.Add(1)
			goModuleListen(ctx, work, wg)
		}
	}
}

//ErrNoHandler4Command 没有处理器
var ErrNoHandler4Command = errors.New("No handler for command")

//HandleTask 处理任务
func HandleTask(cmd Jcommand.Command, t *Jtask.Task) error {
	if workers, ok := handlers[cmd]; ok {
		best := chooseBestWorker(workers)
		if best == nil {
			return ErrNoHandler4Command
		}
		best.taskChannel <- &task{
			cmd:  cmd,
			task: t,
		}
		return nil
	}
	return ErrNoHandler4Command
}

func goRunModule(ctx context.Context, mod Jmodule.Module, wg *sync.WaitGroup) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				Jlog.Error(Jtag.Module(mod.GetName()), "错误关闭", err)
			}
			wg.Done()
		}()
		mod.Run(ctx)
	}()
}

func goModuleListen(ctx context.Context, worker *worker, wg *sync.WaitGroup) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				Jlog.Error(Jtag.Module(worker.mod.GetName()), "监听错误关闭", err)
			}
			wg.Done()
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case task := <-worker.taskChannel:
				worker.handleTask(task)
			}
		}
	}()
}

func chooseBestWorker(ws []*worker) (best *worker) {
	var delay int64 = -1
	best = nil
	for _, w := range ws {
		if delay < 0 || delay > w.delay {
			best = w
		}
	}
	return
}
