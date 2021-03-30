/*
 * @Author: Vongola
 * @FilePath: /JFFun/proto/gen.go
 * @Date: 2021-03-30 11:26:27
 * @Description: file content
 * @描述: 文件描述
 * @LastEditTime: 2021-03-30 16:58:01
 * @LastEditors: Vongola
 */
package proto

//go:generate protoc --proto_path=./ --go_out=../ command.proto errorCode.proto log.proto
