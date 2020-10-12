package main

import (
	"JFFun/gate"
	"JFFun/module"
	"JFFun/server"
	"fmt"
)

func main() {
	fmt.Println("Just For Fun")

	err := server.Run([]module.Module{
		new(gate.M_Gate),
	}...)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("bye")
}
