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
		err2 := processor.unmarshalJSON(buffer)
		if err2 != nil {
			// disconect client
		}
	}
}

func (processor *ServerProcessor) sendMessage(message []byte) {
	processor.connection.Write(message)
}

func (processor *ServerProcessor) unmarshalJSON(j []byte) error {
	var strings map[string]string
	err := json.Unmarshal(j, &strings)
	if err != nil {
		return err
	}
	var typeMessage = strings["type"]
	switch typeMessage {
	case message.IDENTIFY_MESSAGE_TYPE:
		processor.sendMessage(checkIdentify(strings["username"]))
		return nil
	}
	return nil
}




