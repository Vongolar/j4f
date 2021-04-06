package socks5

import (
	"context"
	jconfig "j4f/core/config"
	"j4f/core/task"
	"net"
	"sync"
)

type M_Socks5 struct {
	cfg    config
	listen *net.TCPListener

	wg sync.WaitGroup
}

func (m *M_Socks5) Init(ctx context.Context, name string, cfgPath string) error {
	if err := jconfig.ParseFile(cfgPath, &m.cfg); err != nil {
		return err
	}
	addr, err := net.ResolveTCPAddr("tcp", m.cfg.Address)
	if err != nil {
		return err
	}

	if m.listen, err = net.ListenTCP(addr.Network(), addr); err != nil {
		return err
	}

	go m.listenTCP()
	return nil
}

func (m *M_Socks5) Run(c chan *task.Task) {
LOOP:
	for {
		select {
		case t := <-c:
			if t == nil {
				m.listen.Close()
				break LOOP
			}
		}
	}
}

func (m *M_Socks5) listenTCP() {
	for {
		_, err := m.listen.Accept()
		if err == net.ErrClosed {
			break
		}
		if err != nil {
			continue
		}
	}
}
