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
var invalid bool

// reads sent messages by client
func (processor *ServerProcessor) readMessages() {
	for {
		buffer := make([]byte, 1024)
		length, err1 := processor.connection.Read(buffer)
		if err1 != nil {
			fmt.Println("Error while reading:", err1.Error())	
		}
		messageRecieved, err2 := processor.unmarshalJSON(buffer[:length])
		if err2 != nil {
			processor.sendMessage([]byte("Sorry, bad operation...\n"))
			if identified {
				disconn := message.DisconnectedMessage{message.DISCONNECTED_MESSAGE_TYPE, processor.username}
				toAllUsers(disconn.GetJSON())
			}
			log.Printf("Error: %v\n", err2.Error())
			processor.disconnectClient()
			break
		}
		processor.processMessage(messageRecieved)
		fmt.Printf("message received: %v by %v\n", messageRecieved, processor.username)
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
		counter.blocker.Lock()
		delete(counter.users, processor.username)
		counter.blocker.Unlock()
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
		return message, err
	}
	return message, nil
}

func (processor *ServerProcessor) changeStatus(newStatus string) ([]byte) {
	if accepted := verifyStatus(newStatus); !accepted {
		mess := message.ErrorMessageStatus{message.ERROR_MESSAGE_TYPE, "Invalid status!", message.STATUS_MESSAGE_TYPE}
		invalid = true
		return mess.GetJSON()
	}
	mess := message.SuccesMessage{message.INFO_MESSAGE_TYPE, "status changed succesfully", message.STATUS_MESSAGE_TYPE}
	messUsers := message.NewStatusMessage{message.NEW_STATUS_MESSAGE_TYPE, processor.username, newStatus}
	toAllUsers(messUsers.GetJSON())
	processor.setStatus(newStatus)
	return mess.GetJSON()
}


//processes received messages proceeding in each case as necessary.
func (processor *ServerProcessor) processMessage(gottenMessage map[string]string) {
	var typeMessage string = gottenMessage["type"]
	switch typeMessage {
	case message.IDENTIFY_MESSAGE_TYPE:
		processor.sendMessage(checkIdentify(gottenMessage["username"], processor))
		break
	case message.STATUS_MESSAGE_TYPE:
		processor.sendMessage(processor.changeStatus(gottenMessage["status"]))
		if invalid {
			processor.disconnectClient()
		}
		break
	case message.USERS_MESSAGE_TYPE:
		processor.sendMessage(getUserList())
		break
	case message.DISCONNECT_MESSAGE_TYPE:
		processor.disconnectClient()
		break
	case message.PUBLIC_MESSAGE_TYPE:
		publicMessage := message.NewMessage{message.PUBLIC_MESSAGE_FROM_TYPE, processor.username, gottenMessage["message"]}
		toAllUsers(publicMessage.GetJSON())
		break
	case message.MESSAGE_TYPE:
		err := sendPrivateMessage(gottenMessage["username"], gottenMessage["message"], processor.username)
		if err != nil {
			warning := message.WarningMessageUsername{message.WARNING_MESSAGE_TYPE, "the user received does not exist", message.MESSAGE_TYPE, gottenMessage["username"]}
			processor.sendMessage(warning.GetJSON())
		}
		break
	case message.NEW_ROOM_MESSAGE_TYPE:
		message, _ := createNewRoom(processor.username, processor, gottenMessage["roomname"])
		processor.sendMessage(message)
		break
	case message.INVITE_MESSAGE_TYPE:
		processor.sendMessage(inviteUsersToRoom(gottenMessage["roomname"], gottenMessage["usernames"]))
		break
	case message.JOIN_ROOM_MESSAGE_TYPE:
		processor.sendMessage(joinRoom(processor.username, gottenMessage["roomname"]))
		break
	case message.ROOM_USERS_MESSAGE_TYPE:
		processor.sendMessage(getRoomUserList(processor.username, gottenMessage["roomname"]))
		break
	}
}



