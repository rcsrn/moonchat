package client

import (
	"net"
)

type ClientProcessor struct {
	connection net.Conn
	identified bool
	username string
	messageCreator *messageCreator
}

func getClientInstance(connection net.Conn) *ClientProcessor {
	return &ClientProcessor{connection, false, "", &messageCreator{}}
}

func (processor *ClientProcessor) setConnection(connection net.Conn) {
	processor.connection = connection
}

func (processor *ClientProcessor) sendMessage(message []byte) {
	processor.connection.Write(message)
}

func (processor *ClientProcessor) ProcessMessage(message []string) {
	firstWord := message[0]
	switch firstWord {
	case CLOSE: closeCase(processor)
	}
}

func closeCase(processor *ClientProcessor) {
	disconnectMessage := processor.messageCreator.getDisconnectMessage()
	processor.sendMessage(disconnectMessage)
}

