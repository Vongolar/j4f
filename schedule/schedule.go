package schedule

import (
	"JFFun/config"
	"JFFun/jlog"
	"JFFun/module"
	"JFFun/task"
	"context"
	"fmt"
	"sync"
)

type Schedule struct {
	modules map[string]*mod
}

type mod struct {
	m       module.Module
	channel chan *task.Task
	cfg     moduleConfig
}

func (s *Schedule) RegistModule(m module.Module, name string, cfgPath string) error {
	cfg := new(moduleConfig)
	if err := config.ParseFileConfig(cfgPath, cfg); err != nil {
		return err
	}

	if s.modules == nil {
		s.modules = make(map[string]*mod)
	}

	s.modules[name] = &mod{
		m:       m,
		cfg:     *cfg,
		channel: make(chan *task.Task, cfg.Buffer),
	}

	return nil
}

func (s *Schedule) Run(ctx context.Context) {
	var wg sync.WaitGroup

	for k, m := range s.modules {
		s.goListen(ctx, &wg, k, m)
	}

	for k, m := range s.modules {
		s.goRunModule(ctx, &wg, k, m)
	}
	wg.Wait()
}

func (s *Schedule) goListen(ctx context.Context, wg *sync.WaitGroup, name string, mod *mod) {
	wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				jlog.Error(fmt.Sprintf("module %s handler panic : %v", name, err))

				if mod.cfg.AutoRestart {
					s.goListen(ctx, wg, name, mod)
				}
			}
			wg.Done()
		}()

		for {
			select {
			case <-ctx.Done():
				return

			}
		}

	}()
}

func (s *Schedule) goRunModule(ctx context.Context, wg *sync.WaitGroup, name string, mod *mod) {
	wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				jlog.Error(fmt.Sprintf("module %s panic : %v", name, err))

				if mod.cfg.AutoRestart {
					s.goRunModule(ctx, wg, name, mod)
				}
			}
			wg.Done()
		}()
		mod.m.Run(ctx)
	}()
}
