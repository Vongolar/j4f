package task

import (
	"fmt"
	"testing"

	Jerror "JFFun/data/error"
)

func Test_ChannelTask(t *testing.T) {
	channel := make(chan *Task, 10)
	go func() {
		for {
			select {
			case task := <-channel:
				task.Reply(Jerror.Error_ok, nil)
			}
		}
	}()

	req := NewChannelRequest()
	task1 := &Task{
		Request: req,
	}
	channel <- task1
	resp := req.Wait()
	fmt.Println(resp.ErrCode)
}
