package task

type Task struct {
	CommandID int
	Data      interface{}
}

func (t *Task) Reply(err error, data interface{}) {

}

func (t *Task) ReplyOK(data interface{}) {
	t.Reply(nil, data)
}
