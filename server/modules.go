package server

import (
	Jgate "JFFun/gate"
	Jmodule "JFFun/module"
	Jschedule "JFFun/schedule"
	"io/ioutil"
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

func initModule(m Jmodule.Module, file string) error {
	b, err := ioutil.ReadFile(configPath + file)
	if err != nil {
		return err
	}
	err = m.Init(b)
	if err != nil {
		return err
	}
	Jschedule.Regist(m)
	return nil
}
