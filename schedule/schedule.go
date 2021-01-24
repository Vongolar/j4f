/*
 * @Author: Vongola
 * @FilePath: /JFFun/schedule/schedule.go
 * @Date: 2021-01-24 00:03:02
 * @Description: file content
 * @描述: 文件描述
 * @LastEditTime: 2021-01-24 00:03:41
 * @LastEditors: Vongola
 */

package schedule

import (
	"JFFun/task"
)

type Schedule interface {
	Execute(task *task.Task)
}
