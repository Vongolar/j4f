package main

import (
	"JFFun/gate"
	"JFFun/jlog"
	"JFFun/module"
	"JFFun/server"
)

func main() {
	jlog.Info("Just For Fun")
	server.Run(
		[]module.Module{new(gate.M_Gate), new(gate.M_Gate), new(gate.M_Gate)},
		[]module.Module{new(gate.M_Gate), new(gate.M_Gate)},
		[]module.Module{new(gate.M_Gate)},
	)
	jlog.Info("Good Bye")
}

//go:generate go generate ./proto
