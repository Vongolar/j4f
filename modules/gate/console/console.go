package console

import (
	"bufio"
	"context"
	"io"
	"j4f/core/serialize/json"
	"j4f/core/server"
	"j4f/core/task"
	"j4f/define"
	"os"
	"sync"

	dcommon "j4f/data/common"
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
	req := new(dcommon.Pack)
	err := json.Unmarshal(data, req)
	if err != nil {
		server.ErrTag(m.tag, `消息解析错误`, err)
		return
	}

	if define.IsInvaildCMD(req.GetCommand()) {
		server.ErrTag(m.tag, `无效的消息命令`)
		return
	}

	reqData := define.GetRequestStructByCMD(req.GetCommand())

	if reqData == nil {

	}
}
