/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 18:11:03
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\modules\http\http_test.go
 * @Date: 2021-02-08 17:31:50
 * @描述: 文件描述
 */
package http

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

func Test_Ping(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r1, _ := http.NewRequest(http.MethodPost, "http://localhost:1000", nil)
			r1.Header.Add("Command", "0")
			if resp, err := http.DefaultClient.Do(r1); err == nil {
				fmt.Println(resp.StatusCode)
			} else {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
}
