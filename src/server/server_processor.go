package main

import (
	"net"
	"fmt"
	"encoding/json"
	"github.com/rcsrn/moonchat/src/message"
	"strings"
)

type ServerProcessor struct {
	connection net.Conn
	username string
	status string
}

var identified bool

// reads sent messages by client
func (processor *ServerProcessor) readMessages() {
	for {	
		buffer := make([]byte, 1024)
		_, err1 := processor.connection.Read(buffer)
		if err1 != nil {
			fmt.Println("Error while reading:", err1.Error())	
		}
		messageRecieved, err2 := processor.unmarshalJSON(buffer)
		if err2 != nil {
			processor.sendMessage([]byte("Sorry, bad operation :\n"))
			if identified {
				disconn := message.DisconnectedMessage{message.DISCONNECTED_MESSAGE_TYPE, processor.username}
				toAllUsers(message.GetDisconnectedMessageJSON(disconn))
			} 
			processor.disconnectClient()
			break
		}
		processor.processMessage(messageRecieved)
	}
}

//sets the username to the client once it has been identified.
func (processor *ServerProcessor) setUserName(name string) {
	processor.username = name
	identified = true
}


//disconnects the client when it sends a message out the protocol.
func (processor *ServerProcessor) disconnectClient() {
	processor.connection.Close()
	if value := strings.Compare(processor.username, ""); value != 0 {
		counter.RLock()
		delete(counter.users, processor.username)
		counter.RUnlock()
	}
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

func (processor *ServerProcessor) changeStatus(newStatus string) ([]byte) {
	// if accepted := message.VerifyStatus(newStatus); !accepted {
	// 	mess := message.WarningMessageStatus{message.WARNING_MESSAGE_TYPE, "Invalid status", message.STATUS_MESSAGE_TYPE, newStatus}
	// 	return mess.getJSON()
	// }
	// mess := message.StatusMessage(message.STATUS_MESSAGE_TYPE, newStatus)
	// return mess.getJSON()
	return nil
}



//processes received messages.
func (processor *ServerProcessor)processMessage(messageGotten map[string]string) {
	var typeMessage string = messageGotten["type"]
	switch typeMessage {
	case message.IDENTIFY_MESSAGE_TYPE:
		processor.sendMessage(checkIdentify(messageGotten["username"], processor))
	case message.STATUS_MESSAGE_TYPE:
		processor.sendMessage(processor.changeStatus(messageGotten["status"]))
	}
	// other cases must be implemented.
}



