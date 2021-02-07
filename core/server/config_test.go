/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-05 10:17:17
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\server\config_test.go
 * @Date: 2021-02-04 14:30:13
 * @描述: 文件描述
 */

package server

import (
	"fmt"
	"testing"
)

func Example_ParseFlagConfigPaths() {
	paths := parseConfigs("1;2;3;;")
	fmt.Println(paths)
	//Output: [1 2 3]
}

func Test_CloseChannel(t *testing.T) {
	c := make(chan int)
	close(c)
	close(c)
	fmt.Println(<-c)
}
