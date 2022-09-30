package main

import (
	"j4f/server"
	"log"
)

func main() {
	server := new(server.Server)
	err := server.Run()
	if err != nil {
		log.Println(err)
	}
}
