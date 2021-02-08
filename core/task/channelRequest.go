/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 18:01:48
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\core\task\channelRequest.go
 * @Date: 2021-02-08 17:52:27
 * @描述: 文件描述
 */

package task

import (
	"io"
	"j4f/data"
)

func NewChannelRequest(w io.Writer) *ChannelRequest {
	return &ChannelRequest{
		w: w,
		c: make(chan data.Error, 1),
	}
}

type ChannelRequest struct {
	w io.Writer
	c chan data.Error
}

func (r *ChannelRequest) GetWriter() io.Writer {
	return r.w
}

func (r *ChannelRequest) Reply(err data.Error) {
	r.c <- err
	close(r.c)
}

func (r *ChannelRequest) Wait() data.Error {
	return <-r.c
}
