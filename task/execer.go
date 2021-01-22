/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-22 19:02:58
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\task\execer.go
 * @Date: 2021-01-22 18:52:19
 * @描述: 文件描述
 */
package task

type Execer interface {
	Execute()
	ExecuteLocal(*Task)
	Distribute()
}
