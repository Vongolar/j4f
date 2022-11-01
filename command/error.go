package command

import "errors"

var Err_MODULE_CLOSE = errors.New("服务器关闭，不接受处理")
var Err_TASK_DATA_INVAILD = errors.New("任务类型不匹配")
var Err_REGIST_CONFLICT_TASK_COMMAND = errors.New("不同类型模块注册任务冲突")
