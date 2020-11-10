package jtag

import (
	"fmt"
)

//Server 服务器
var Server = `server`

//Schedule 调度器
var Schedule = `schedule`

//Gate 网关模块
var Gate = `gate`

//Login 登陆模块
var Login = `login`

//Match 匹配模块
var Match = `match`

//Game 游戏
func Game(id string) string {
	return fmt.Sprintf("room-%s", id)
}

//Net 网络
var Net = `net`

//DataBase 数据库
var DataBase = `db`

//Cache 缓存数据库
var Cache = `cache`
