package jschedule

import (
	"JFFun/data/Dcommand"
	"JFFun/data/Derror"
	jtask "JFFun/task"
	"context"
	"math"
	"sync"
	"time"
)

//HandleTask 自动选取一个最优的模块处理消息
func HandleTask(cmd Dcommand.Command, t *jtask.Task) {
	if mods, exist := handlerRoute[cmd]; exist && len(mods) > 0 {
		var fastMod *module
		minDelay := time.Duration(math.MaxInt64)
		for _, mod := range mods {
			if len(mod.taskChannel) == 0 && mod.isAvailability() {
				mod.taskChannel <- &task{
					cmd:  cmd,
					task: t,
				}
				return
			}

			if mod.isAvailability() && mod.getDelay() < minDelay {
				minDelay = mod.getDelay()
				fastMod = mod
			}
		}
		fastMod.taskChannel <- &task{
			cmd:  cmd,
			task: t,
		}
		return
	}
	t.Error(Derror.Error_noHandler)
}

//HandleTaskBy 指定模块
func HandleTaskBy(module string, cmd Dcommand.Command, t *jtask.Task) {
	if mods, exist := handlerRoute[cmd]; exist && len(mods) > 0 {
		for _, mod := range mods {
			if mod.name == module {
				mod.taskChannel <- &task{
					cmd:  cmd,
					task: t,
				}
				return
			}
		}
	}
	t.Error(Derror.Error_noHandler)
}

//HandleTaskUntilOK 顺序执行直到一个正确返回,如果没有正确的返回最后一个结果
func HandleTaskUntilOK(cmd Dcommand.Command, t *jtask.Task) {
	if mods, exist := handlerRoute[cmd]; exist && len(mods) > 0 {
		for i, mod := range mods {
			creq := jtask.NewInnerRequest()
			ctask := &jtask.Task{
				Request:  creq,
				PlayerID: t.PlayerID,
				Data:     t.Data,
				Raw:      t.Raw,
			}
			mod.taskChannel <- &task{
				cmd:  cmd,
				task: ctask,
			}
			if err, resp := creq.Wait(); err == Derror.Error_ok || i+1 == len(mods) {
				t.Reply(err, resp)
				return
			}
		}
	}
	t.Error(Derror.Error_noHandler)
}

//MutliResponse 多模块执行响应
type MutliResponse = map[string]interface{}

//HandleTaskByAll 所有模块顺序执行，只返回正确的结果,注意如果在模块Handler线程内会导致线程卡死
func HandleTaskByAll(cmd Dcommand.Command, t *jtask.Task) {
	HandleTaskByOthers("", cmd, t)
}

//HandleTaskByOthers 除了selfMoudle以外，所有模块顺序执行，只返回正确的结果
func HandleTaskByOthers(selfMoudle string, cmd Dcommand.Command, t *jtask.Task) {
	if mods, exist := handlerRoute[cmd]; exist && len(mods) > 0 {
		res := make(MutliResponse, len(mods))
		for _, mod := range mods {
			if mod.name == selfMoudle {
				continue
			}
			req := jtask.NewInnerRequest()
			st := &jtask.Task{
				Request:  req,
				PlayerID: t.PlayerID,
				Raw:      t.Raw,
				Data:     t.Data,
			}
			mod.taskChannel <- &task{
				cmd:  cmd,
				task: st,
			}
			if err, resp := req.Wait(); err == Derror.Error_ok {
				res[mod.name] = resp
			}
		}
		t.Reply(Derror.Error_ok, res)
		return
	}
	t.Error(Derror.Error_noHandler)
}

//SyncData 通知结构
type SyncData struct {
	Scmd Dcommand.SyncCommand
	Data interface{}
}

//Sync 通知玩家
func Sync(scmd Dcommand.SyncCommand, data interface{}, playerID string) {
	t := &jtask.Task{
		PlayerID: playerID,
		Data: &SyncData{
			Scmd: scmd,
			Data: data,
		},
	}
	if mods, exist := handlerRoute[Dcommand.Command_sync]; exist && len(mods) > 0 {
		for _, mod := range mods {
			mod.taskChannel <- &task{
				cmd:  Dcommand.Command_sync,
				task: t,
			}
		}
	}
}

//Run 运行
func Run(ctx context.Context, wg *sync.WaitGroup) {
	goRunModules(ctx, wg)
}
