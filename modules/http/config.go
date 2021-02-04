/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-04 18:47:47
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\modules\http\config.go
 * @Date: 2021-02-04 18:47:04
 * @描述: 文件描述
 */

package http

type config struct {
	Port int `toml:"port" yaml:"port" json:"port"`
}
