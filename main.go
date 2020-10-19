package main

import (
	Jlog "JFFun/log"
	Jtag "JFFun/log/tag"
	Jserver "JFFun/server"
)

func main() {
	Jlog.Info(Jtag.Server, `Just For Fun`)
	Jserver.Run()
}
