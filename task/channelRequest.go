package task

//NewChannelRequest 创建一个模块间请求
func NewChannelRequest() *ChannelRequest {
	return &ChannelRequest{
		channel: make(chan *ResponseData),
	}
}

//ChannelRequest 程序内部模块间请求
type ChannelRequest struct {
	channel chan *ResponseData
}

//Wait 等待请求回应
func (req *ChannelRequest) Wait() *ResponseData {
	return <-req.channel
}

//Reply 回应
func (req *ChannelRequest) Reply(resp *ResponseData) error {
	req.channel <- resp
	return nil
}
