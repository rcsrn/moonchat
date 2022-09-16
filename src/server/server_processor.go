package main

import (
	"net"
	"fmt"
)

type ServerProcessor struct {
	connection net.Conn
}

func (processor *ServerProcessor) processClient() {
	fmt.Println("Procesando al cliente")
}

func (processor *ServerProcessor) sendToClient(message []byte) {
	processor.connection.Write(message)
}

