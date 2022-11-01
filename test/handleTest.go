package test

import (
	"fmt"
	"j4f/task"
	"time"
)

func (m *Module) testHandle(t *task.Task) {
	fmt.Println("测试")
	time.Sleep(time.Second * 10)
	panic("测试崩溃")
}
