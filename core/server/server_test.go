/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 17:05:00
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\server\server_test.go
 * @Date: 2021-02-04 11:30:52
 * @描述: 文件描述
 */
package server

import (
	"fmt"
	"testing"
)

func Test_ChannelLen(t *testing.T) {
	c := make(chan int, 20)
	fmt.Println(len(c))
	c <- 20
	fmt.Println(len(c))
	fmt.Println(cap(c))
}
