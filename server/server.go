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
	"sync"
)

func Run(modules ...module.Module) {
	RunServers(modules)
}

func RunServers(modules ...[]module.Module) {
	release, cfgs := parseFlag()

	if release {
		jlog.Info(`run in release mode`)
	} else {
		jlog.Info(`run in debug mode`)
	}

	if release && len(modules) > 1 {
		jlog.Warning(`it's better to pack all modules to one server.`)
	}

	if len(modules) > len(cfgs) {
		jlog.Error(`length of args '-cfg' less than server's length`)
		return
	}

	var ss []*server
	for i := range modules {
		cfg := new(config)
		if err := jconfig.ParseLocalConfig(cfgs[i], cfg); err != nil {
			jlog.Error(fmt.Sprintf("parse server config file '%s' error", cfgs[i]), err)
			return
		}

		ss = append(ss, &server{cfg: *cfg})
	}

	ctx := context.Background()
	var cancel context.CancelFunc
	if release {
		ctx, cancel = context.WithCancel(context.Background())
	}

	var mwg sync.WaitGroup
	for i, s := range ss {
		server, mods := s, modules[i]
		server.run(ctx, &mwg, mods...)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	if cancel != nil {
		cancel()
	}

	mwg.Wait()
}

type server struct {
	schedule *schedule.Schedule

	cfg config
}

func (s *server) run(ctx context.Context, wg *sync.WaitGroup, modules ...module.Module) {
	if len(s.cfg.Modules) < len(modules) {
		jlog.Error(fmt.Sprintf("server %s's module configs less than modules.", s.cfg.Name))
		return
	}

	if err := s.initModules(modules); err != nil {
		jlog.Error(fmt.Sprintf("server %s init modules error.", s.cfg.Name), err)
		return
	}

	s.schedule = new(schedule.Schedule)
	for i, m := range modules {
		if err := s.schedule.RegistModule(s.cfg.Modules[i].Name, m, s.cfg.Modules[i].Path); err != nil {
			jlog.Error(fmt.Sprintf("server %s regist module %s error.", s.cfg.Name, s.cfg.Modules[i].Name), err)
			return
		}
	}

	jlog.Info(fmt.Sprintf("server %s init success", s.cfg.Name))

	s.schedule.Run(ctx, wg)
}

func (s *server) initModules(modules []module.Module) error {
	for i, m := range modules {
		if err := m.Init(s.cfg.Modules[i].Name, s.cfg.Modules[i].Path); err != nil {
			return err
		}
	}
	return nil
}
