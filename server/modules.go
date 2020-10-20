package server

import (
	Jgate "JFFun/gate"
	Jmodule "JFFun/module"
	Jschedule "JFFun/schedule"
)

func initModules() error {
	for _, c := range cfg.Gate {
		m := &Jgate.MGate{}
		if err := initModule(m, c); err != nil {
			return err
		}
	}

	return nil
}

func initModule(m Jmodule.Module, cfg string) error {
	if err := m.Init(cfg); err != nil {
		return err
	}
	Jschedule.Regist(m)
	return nil
}
