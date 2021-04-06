package socks5

import (
	"context"
	"j4f/core/task"
	"net"
)

type M_Socks5 struct {
}

func (m *M_Socks5) Init(ctx context.Context, name string, cfgPath string) error {
	return nil
}

func (m *M_Socks5) Run(c chan *task.Task) {
	addr, _ := net.ResolveTCPAddr("tcp", ":9999")
	l, _ := net.ListenTCP(addr.Network(), addr)

	go func() {
		l.Accept()
	}()

LOOP:
	for {
		select {
		case t := <-c:
			if t == nil {
				l.Close()
				break LOOP
			}
		}
	}
}
