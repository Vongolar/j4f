/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-15 11:29:25
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\proto\proto_gen.go
 * @Date: 2021-01-14 10:22:11
 * @描述: 文件描述
 */
package proto

//go:generate protoc --proto_path=./ --go_out=../data game.proto

//go:generate protoc --proto_path=./ --go_out=../data command.proto
