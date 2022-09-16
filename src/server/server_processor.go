package main

import (
	"net"
)

type ServerProcessor struct {
	connection net.Conn
}

func (processor *ServerProcessor) processClient() {
	
}

func (processor *ServerProcessor) sendToClient(message []byte) {
	processor.connection.Write(message)
}

