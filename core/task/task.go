/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-07 19:22:44
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\task\task.go
 * @Date: 2021-02-04 17:44:05
 * @描述: 文件描述
 */
package task

import "j4f/data"

type Task struct {
	CMD     data.Command
	Data    interface{}
	RawData []byte
	NextCMD []data.Command
}

type TaskHandleTuple struct {
}
