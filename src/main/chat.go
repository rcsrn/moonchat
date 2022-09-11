package main

import (
	"fmt"
)

type chat struct {
}

func (chat *chat) start() {
	server := server{}
	listener, err := server.createServer()
	if err != nil {
		fmt.Println("[1] Algo ha salido mal. :(")
	}
	server.waitForConnections(listener)
}


