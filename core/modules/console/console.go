package console

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"j4f/core/request"
	"j4f/core/serialize/json"
	"j4f/core/server"
	"j4f/core/task"
	"j4f/define"
	"os"
	"sync"
	"unsafe"

	dcommon "j4f/data/common"
	"j4f/data/errCode"
)

type M_Console struct {
	tag         string
	stdinReader *bufio.Reader
	wg          sync.WaitGroup
}

func (m *M_Console) Init(ctx context.Context, name string, cfgPath string) error {
	m.tag = name
	m.stdinReader = bufio.NewReader(os.Stdin)
	m.wg.Add(1)
	go m.listen()
	return nil
}

func (m *M_Console) Run(c chan *task.Task) {
LOOP:
	for {
		select {
		case t := <-c:
			if t == nil {
				os.Stdin.Close()
				break LOOP
			}
		}
	}

	m.wg.Wait()
}

func (m *M_Console) listen() {
	var cache []byte
	for {
		line, full, err := m.stdinReader.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil {
			server.ErrTag(m.tag, err)
			continue
		}

		if full {
			cache = append(cache, line...)
			continue
		}

		if len(cache) > 0 {
			line = append(cache, line...)
			cache = nil
		}

		if len(line) == 0 {
			continue
		}

		m.serialize(line)
	}
	m.wg.Done()
}

func (m *M_Console) serialize(data []byte) {
	req := new(dcommon.ConsolePack)
	err := json.Unmarshal(data, req)
	if err != nil {
		server.ErrTag(m.tag, `消息解析错误`, err)
		return
	}

	if !server.EqualConsoleKey(req.GetKey()) {
		server.ErrTag(m.tag, `非法的Console命令`)
		return
	}

	if define.IsInvaildCMD(req.GetCommand()) {
		server.ErrTag(m.tag, `无效的消息命令`)
		return
	}

	reqData := define.GetRequestStructByCMD(req.GetCommand())

	t := &task.Task{CMD: req.GetCommand(), Author: define.Auth_Console}
	if reqData == nil {
		t.Data = req.GetData()
	} else {
		s := req.GetData()
		err = json.Unmarshal(*(*[]byte)(unsafe.Pointer(&s)), reqData)
		if err != nil {
			server.ErrTag(m.tag, `消息解析错误`, err)
			return
		}
		t.Data = reqData
	}
	r := request.CreateSyncRequest()
	t.Request = r

	if err = server.Handle(t); err != nil {
		return
	}

	ec, resp := r.Wait()

	if ec == errCode.Code_ok {
		fmt.Printf("ok : %v\n", resp)
	} else {
		fmt.Printf("error %d : %s\n", ec, errCode.Code_name[int32(ec)])
	}
}
