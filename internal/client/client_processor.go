package client

import (
	"net"
)

type ClientProcessor struct {
	connection net.Conn
	identified bool
	username string
}

func getClientInstance(connection net.Conn) *ClientProcessor {
	return &ClientProcessor{connection, false, ""}
}

func (processor *ClientProcessor) setConnection(connection net.Conn) {
	processor.connection = connection
}
