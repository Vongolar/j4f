package main

import (
	"j4f/server"
)

func main() {
	server := new(server.Server)
	server.Run()
}
