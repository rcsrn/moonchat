package main

import (
	"net"
	"fmt"
	"encoding/json"
	"github.com/rcsrn/moonchat/src/message"
)

type ServerProcessor struct {
	connection net.Conn
}

func (processor *ServerProcessor) readMessages() {
	for {
		buffer := make([]byte, 1024)
		_, err1 := processor.connection.Read(buffer)
		if err1 != nil {
			fmt.Println("Error reading:", err1.Error())
		}
		message, err2 := processor.unmarshalJSON(buffer)
		if err2 != nil {
			// disconect client
		}
		processor.processMessage(message)
	}
}

func (processor *ServerProcessor) sendMessage(message []byte) {
	processor.connection.Write(message)
}

func (processor *ServerProcessor) unmarshalJSON(j []byte) (map[string]string, error) {
	var message map[string]string
	err := json.Unmarshal(j, &message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (processor *ServerProcessor)processMessage(messageGotten map[string]string) {
	var typeMessage = messageGotten["type"]
	switch typeMessage {
	case message.IDENTIFY_MESSAGE_TYPE:
		processor.sendMessage(checkIdentify(messageGotten["username"], processor))
	}
	//Faltan todos los demas casos
}


