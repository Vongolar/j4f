package server

import (
	jconfig "JFFun/config"
	"JFFun/jlog"
	"JFFun/module"
	"JFFun/schedule"
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
)

func Run(modules ...module.Module) {
	RunServers(modules)
}

func RunServers(modules ...[]module.Module) {
	release, cfgs := parseFlag()

	if release && len(modules) > 1 {
		jlog.Warning(`it's better to pack all modules to one server.`)
	}

	if len(modules) > len(cfgs) {
		jlog.Error(`length of args '-cfg' less than server's length`)
		return
	}

	var ss []*server
	for i, _ := range modules {
		cfg := new(config)
		if err := jconfig.ParseFileConfig(cfgs[i], cfg); err != nil {
			jlog.Error(fmt.Sprintf("parse config file '%s' error", cfgs[i]), err)
			return
		}

		dir, _ := filepath.Split(cfgs[i])
		if err := cfg.checkModuleConfigExist(dir); err != nil {
			return
		}

		ss = append(ss, &server{name: cfg.Name, cfg: *cfg})
	}

	ctx := context.Background()
	var cancel context.CancelFunc
	if release {
		ctx, cancel = context.WithCancel(context.Background())
	}

	var wg sync.WaitGroup
	for i, s := range ss {
		server, mods := s, modules[i]
		wg.Add(1)
		go func() {
			server.run(ctx, mods...)
			wg.Done()
		}()
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	cancel()

	wg.Wait()
}

type server struct {
	schedule *schedule.Schedule

	name string
	cfg  config
}

func (s *server) run(ctx context.Context, modules ...module.Module) {
	if len(s.cfg.Modules) < len(modules) {
		jlog.Error(fmt.Sprintf("server %s's module configs less than modules.", s.name))
		return
	}

	if err := s.initModules(modules); err != nil {
		jlog.Error(fmt.Sprintf("server %s init modules error.", s.name), err)
		return
	}

	s.schedule = new(schedule.Schedule)
	for i, m := range modules {
		if err := s.schedule.RegistModule(m, s.cfg.Modules[i].Name, s.cfg.Modules[i].Path); err != nil {
			jlog.Error(fmt.Sprintf("server %s regist module %s error.", s.name, s.cfg.Modules[i].Name), err)
			return
		}
	}

	jlog.Info(fmt.Sprintf("server %s init success", s.name))

	s.schedule.Run(ctx)
}

func (s *server) initModules(modules []module.Module) error {
	for i, m := range modules {
		if err := m.Init(s.cfg.Modules[i].Name, s.cfg.Modules[i].Path); err != nil {
			return err
		}
	}
	return nil
}
