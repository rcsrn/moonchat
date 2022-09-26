package main

import (
	"net"
	"fmt"
	"encoding/json"
	"github.com/rcsrn/moonchat/src/message"
)

type ServerProcessor struct {
	connection net.Conn
	user string
}

// reads messages sent by client
func (processor *ServerProcessor) readMessages() {
	for {	
		buffer := make([]byte, 1024)
		_, err1 := processor.connection.Read(buffer)
		if err1 != nil {
			fmt.Println("Error while reading:", err1.Error())
			
		}
		message, err2 := processor.unmarshalJSON(buffer)
		if err2 != nil {
			processor.disconnectClient()
		}
		processor.processMessage(message)
	}
}

//sets the username to the client once it has been identified.
func (processor *ServerProcessor) setUserName(userName string) {
	processor.user = userName
}


//disconnects the client when it sends a message out the protocol.
func (processor *ServerProcessor) disconnectClient() {
	processor.connection.Close()
	counter.RLock()
	delete(counter.users, processor.user)
	counter.RUnlock()
}

//sends messages to the client through its connection.
func (processor *ServerProcessor) sendMessage(message []byte) {
	if processor.connection != nil {
		processor.connection.Write(message)
	}
}

//unmarshals messages that need to be processed by serverProcessor
func (processor *ServerProcessor) unmarshalJSON(j []byte) (map[string]string, error) {
	var message map[string]string
	err := json.Unmarshal(j, &message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

//processes received messages.
func (processor *ServerProcessor)processMessage(messageGotten map[string]string) {
	var typeMessage string = messageGotten["type"]
	switch typeMessage {
	case message.IDENTIFY_MESSAGE_TYPE:
		processor.sendMessage(checkIdentify(messageGotten["username"], processor))
	}
	// other cases must be implemented.
}



