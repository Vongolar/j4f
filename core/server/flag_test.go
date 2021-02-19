/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-19 14:34:18
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\server\flag_test.go
 * @Date: 2021-02-19 14:26:31
 * @描述: 文件描述
 */

package server

import "fmt"

func Example_ParseConfigFiles() {
	fmt.Println(parseConfigFiles("server1.toml; ;server2.toml;"))
	//Output: [server1.toml server2.toml]
}
