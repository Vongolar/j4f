/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 17:21:28
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\author\author.go
 * @Date: 2021-02-08 16:25:52
 * @描述: 文件描述
 */

package author

import (
	"j4f/data"
)

type Author = int

const (
	None   = 1  //00000001
	Guest  = 2  //00000010
	User   = 4  //00000100
	Server = 16 //00001000
)
const all = None | Guest | User | Server

func Auth(cmd data.Command, auth Author) bool {
	return getAuthByCmd(cmd)&auth != 0
}

func getAuthByCmd(cmd data.Command) int {
	switch cmd {
	case data.Command_ping:
		return all
	}
	return 0
}
