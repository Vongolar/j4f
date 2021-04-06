package mhttp

import (
	"context"
	jconfig "j4f/core/config"
	"j4f/core/task"
	"net/http"
	"os"
)

type M_Http struct {
	cfg config

	httpServer *http.Server
	mux        *http.ServeMux
}

func (m *M_Http) Init(ctx context.Context, name string, cfgPath string) error {
	err := jconfig.ParseFile(cfgPath, &m.cfg)
	if err != nil {
		return err
	}

	m.httpServer = &http.Server{Addr: m.cfg.Address, Handler: m.getHttpMux()}

	if len(m.cfg.Cert) == 0 && len(m.cfg.Key) == 0 {
		go m.httpServer.ListenAndServe()
		return nil
	}

	_, cerr := os.Stat(m.cfg.Cert)
	_, kerr := os.Stat(m.cfg.Key)

	if os.IsExist(cerr) && os.IsExist(kerr) {
		go m.httpServer.ListenAndServeTLS(m.cfg.Cert, m.cfg.Key)
		return nil
	}

	return os.ErrNotExist
}

func (m *M_Http) Run(c chan *task.Task) {
LOOP:
	for {
		select {
		case t := <-c:
			if t == nil {
				m.httpServer.Close()
				break LOOP
			}
		}
	}
}
