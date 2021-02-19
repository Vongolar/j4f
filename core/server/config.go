/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-19 15:45:26
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\server\config.go
 * @Date: 2021-02-19 15:42:41
 * @描述: 文件描述
 */

package server

type config struct {
	Name string `json:"name" toml:"name" yaml:"name"`
}
