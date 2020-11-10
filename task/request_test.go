package jtask

import (
	"JFFun/data/Derror"
	"fmt"
	"testing"
	"time"
)

func Test_InnerChannel(t *testing.T) {
	req := NewInnerRequest()
	go func() {
		time.Sleep(time.Second * 4)
		req.Reply(Derror.Error_ok, 2)
	}()

	fmt.Println(req.Wait())
}

func Test_InnerChannelInOneChannel(t *testing.T) {
	req := NewInnerRequest()
	req.Reply(Derror.Error_ok, 2)
	fmt.Println(req.Wait())
}
