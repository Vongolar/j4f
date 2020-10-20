package server

import (
	Jgate "JFFun/gate"
	Jmodule "JFFun/module"
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
	return m.Init(cfg)
}
