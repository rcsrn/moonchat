package main

import (
	"net"
	"fmt"
	"encoding/json"
	"github.com/rcsrn/moonchat/src/message"
	"strings"
	"log"
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
			processor.sendMessage([]byte("Sorry, bad operation...\n"))
			if identified {
				disconn := message.DisconnectedMessage{message.DISCONNECTED_MESSAGE_TYPE, processor.username}
				toAllUsers(disconn.GetJSON())
			}
			log.Printf("Error: %v\n", err2)
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

func (processor *ServerProcessor) setStatus(newStatus string) {
	processor.status = newStatus
}

//disconnects the client when it sends a message out the protocol.
func (processor *ServerProcessor) disconnectClient() {
	processor.connection.Close()
	if value := strings.Compare(processor.username, ""); value != 0 {
		counter.RLock()
		delete(counter.users, processor.username)
		counter.RUnlock()
		mess := message.DisconnectedMessage{message.DISCONNECTED_MESSAGE_TYPE, processor.username}
		toAllUsers(mess.GetJSON())
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
	if accepted := verifyStatus(newStatus); !accepted {
		mess := message.WarningMessageStatus{message.WARNING_MESSAGE_TYPE, "Invalid status", message.STATUS_MESSAGE_TYPE, newStatus}
		return mess.GetJSON()
	}
	mess := message.SuccesMessage{message.INFO_MESSAGE_TYPE, "status changed succesfully", message.STATUS_MESSAGE_TYPE}
	messUsers := message.NewStatusMessage{message.NEW_STATUS_MESSAGE_TYPE, processor.username, newStatus}
	toAllUsers(messUsers.GetJSON())
	processor.setStatus(newStatus)
	return mess.GetJSON()
}


//processes received messages.
func (processor *ServerProcessor)processMessage(messageGotten map[string]string) {
	var typeMessage string = messageGotten["type"]
	switch typeMessage {
	case message.IDENTIFY_MESSAGE_TYPE:
		processor.sendMessage(checkIdentify(messageGotten["username"], processor))
		break
	case message.STATUS_MESSAGE_TYPE:
		processor.sendMessage(processor.changeStatus(messageGotten["status"]))
		break
	case message.USERS_MESSAGE_TYPE:
		processor.sendMessage(getUserList())
		break
	case message.DISCONNECT_MESSAGE_TYPE:
		processor.disconnectClient()
		break
		
	}
	// other cases must be implemented.
}



