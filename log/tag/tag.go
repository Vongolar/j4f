package tag

import (
	"fmt"
)

//Server 服务器
var Server = `[server]`

//Schedule 调度器
var Schedule = `[shedule]`

//Module 模块
func Module(mod string) string {
	return fmt.Sprintf("[mod %s]", mod)
}

//Net 网络
var Net = `[net]`
