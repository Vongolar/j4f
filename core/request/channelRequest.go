package request

import "j4f/data/errCode"

type responseData struct {
	code errCode.Code
	ext  interface{}
}

func CreateSyncRequest() *channelRequest {
	return &channelRequest{c: make(chan *responseData)}
}

func CreateAsyncRequest() *channelRequest {
	return &channelRequest{c: make(chan *responseData, 1)}
}

type channelRequest struct {
	c chan *responseData
}

func (r *channelRequest) Reply(code errCode.Code, ext interface{}) {
	r.c <- &responseData{code: code, ext: ext}
	close(r.c)
}

func (r *channelRequest) Wait() (errCode.Code, interface{}) {
	resp := <-r.c
	return resp.code, resp.ext
}
