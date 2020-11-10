package main

import (
	jgate "JFFun/gate"
	jlog "JFFun/log"
	jlogin "JFFun/login"
	jmatch "JFFun/match"
	jmodule "JFFun/module"
	jroom "JFFun/room"
	jserver "JFFun/server"
)

func main() {
	jlog.Info(``, `Just For Fun`)

	jserver.Regist(`gate`, func() jmodule.Module { return new(jgate.MGate) })
	jserver.Regist(`login`, func() jmodule.Module { return new(jlogin.MLogin) })
	jserver.Regist(`match`, func() jmodule.Module { return new(jmatch.MMatch) })
	jserver.Regist(`room`, func() jmodule.Module { return new(jroom.MRoom) })

	jserver.Run()
}
