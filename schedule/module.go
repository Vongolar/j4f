package jschedule

import (
	"JFFun/data/Dcommand"
	"JFFun/data/Derror"
	jlog "JFFun/log"
	jtag "JFFun/log/tag"
	jmodule "JFFun/module"
	jtask "JFFun/task"
	"context"
	"fmt"
	"sync"
	"time"
)

var taskTimeout = time.Minute
var moduleRestartTimeout = time.Minute * 10

type module struct {
	handlerTaskCnt   int           //近期处理任务数
	perHandlerDur    time.Duration //近期处理任务平均时间
	lastTaskTime     time.Time     //最近任务开始时间
	lastTaskComplete bool

	name        string
	mod         jmodule.Module
	taskChannel chan *task
	handlers    map[Dcommand.Command]func(*jtask.Task)
}

var modules []*module
var handlerRoute = make(map[Dcommand.Command][]*module)

//RegistModule 注册模块
func RegistModule(name string, mod jmodule.Module, buf int) {
	m := &module{
		name:        name,
		mod:         mod,
		taskChannel: make(chan *task, buf),
		handlers:    mod.GetHandler(),
	}

	for k := range m.handlers {
		handlerRoute[k] = append(handlerRoute[k], m)
	}
	modules = append(modules, m)

	jlog.Info(jtag.Schedule, fmt.Sprintf("注册模块: %s", m.name))
}

func goRunModules(ctx context.Context, wg *sync.WaitGroup) {
	for _, mod := range modules {
		wg.Add(2)
		go mod.listenCommand(ctx, wg)
		go mod.run(ctx, wg)
	}
}

type task struct {
	cmd  Dcommand.Command
	task *jtask.Task
}

func (mod *module) listenCommand(ctx context.Context, wg *sync.WaitGroup) {
	mod.reset()
	defer func() {
		if err := recover(); err != nil {
			jlog.Warning(jtag.Schedule, fmt.Sprintf("%s模块意外关闭监听", mod.name), err)
			jlog.Info(jtag.Schedule, fmt.Sprintf("%s模块重新监听", mod.name))
			go mod.listenCommand(ctx, wg)
			return
		}
		wg.Done()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case task := <-mod.taskChannel:
			if handler, exist := mod.handlers[task.cmd]; exist {
				mod.onTaskStart()
				handler(task.task)
				mod.onTaskEnd()
			} else {
				task.task.Error(Derror.Error_noHandler)
			}
		}
	}
}

func (mod *module) run(ctx context.Context, wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			jlog.Warning(jtag.Schedule, fmt.Sprintf("%s模块意外停止运行", mod.name), err)
			jlog.Info(jtag.Schedule, fmt.Sprintf("%s模块重新运行", mod.name))
			go mod.run(ctx, wg)
			return
		}
		wg.Done()
	}()
	mod.mod.Run(ctx, mod.name)
}

func (mod *module) onTaskStart() {
	mod.lastTaskTime = time.Now()
	mod.lastTaskComplete = false
}

func (mod *module) onTaskEnd() {
	mod.lastTaskComplete = true
	dur := time.Now().Sub(mod.lastTaskTime)
	allDur := time.Duration(mod.handlerTaskCnt)*mod.perHandlerDur + dur
	mod.handlerTaskCnt++
	mod.perHandlerDur = allDur / time.Duration(mod.handlerTaskCnt)

	if mod.handlerTaskCnt > 10000 {
		mod.handlerTaskCnt >>= 1
	}
}

func (mod *module) isAvailability() bool {
	if mod.lastTaskComplete {
		return true
	}

	if mod.lastTaskTime.Add(time.Minute).Before(time.Now()) {
		return false
	}
	return true
}

func (mod *module) getDelay() time.Duration {
	if mod.isAvailability() {
		return mod.perHandlerDur
	}
	return -1
}

func (mod *module) reset() {
	mod.handlerTaskCnt = 0
	mod.perHandlerDur = 0
	mod.lastTaskComplete = true
}
