/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-07 19:30:07
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\gen\gen.go
 * @Date: 2021-02-04 16:39:29
 * @描述: 文件描述
 */

package main

//生成proto
//go:generate protoc --proto_path=../proto/ --go_out=../data command.proto
//go:generate protoc --proto_path=../proto/ --go_out=../data login/login.proto
