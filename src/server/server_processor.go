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
	userStatus string
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
				toAllUsers(processor, disconn.GetJSON())
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

func (processor *ServerProcessor) setUserStatus(newStatus string) {
	processor.userStatus = newStatus
}

//disconnects the client when it sends a message out the protocol.
func (processor *ServerProcessor) disconnectClient() {
	processor.connection.Close()
	if value := strings.Compare(processor.username, ""); value != 0 {
		counter.blocker.Lock()
		delete(counter.users, processor.username)
		counter.blocker.Unlock()
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

func (processor *ServerProcessor) changeStatus(newStatus string) (bool) {
	if accepted := verifyStatus(newStatus); !accepted {
		return false
	}
	processor.setUserStatus(newStatus)
	return true
}


//processes received messages proceeding in each case as necessary.
func (processor *ServerProcessor) processMessage(gottenMessage map[string]string) {
	var typeMessage string = gottenMessage["type"]
	switch typeMessage {
	case message.IDENTIFY_MESSAGE_TYPE:
		identifyCase(processor, gottenMessage["username"])
		break
	case message.STATUS_MESSAGE_TYPE:
		statusCase(processor, gottenMessage["status"])
		break
	case message.USERS_MESSAGE_TYPE:
		userListCase(processor)
		break
	case message.DISCONNECT_MESSAGE_TYPE:
		disconnectClientCase(processor)
		break
	case message.PUBLIC_MESSAGE_TYPE:
		publicMessageCase(processor, gottenMessage["message"])
		break
	case message.MESSAGE_TYPE:
		privateMessageCase(processor, gottenMessage["username"], gottenMessage["message"])
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

func statusCase(processor *ServerProcessor, newStatus string) {
	if value := strings.Compare(processor.userStatus, newStatus); value == 0 {
		str := fmt.Sprintf("El estado actual es '%s'. Favor de elegir otro.", processor.userStatus)
		warning := message.WarningMessageStatus{message.WARNING_MESSAGE_TYPE, str, message.STATUS_MESSAGE_TYPE, processor.userStatus}
		processor.sendMessage(warning.GetJSON())
		return
	}
	
	changedStatus := processor.changeStatus(newStatus)
	if changedStatus {
		succes := message.SuccesMessage{message.INFO_MESSAGE_TYPE, "Succes: status changed succesfully", message.STATUS_MESSAGE_TYPE}
		processor.sendMessage(succes.GetJSON())
		messageToUsers := message.NewStatusMessage{message.NEW_STATUS_MESSAGE_TYPE, processor.username, newStatus}
		toAllUsers(processor, messageToUsers.GetJSON())
	} else {
		error := message.ErrorMessageStatus{message.ERROR_MESSAGE_TYPE, "Invalid status!", message.STATUS_MESSAGE_TYPE}
		processor.sendMessage(error.GetJSON())
		processor.disconnectClient()
		messageToUsers := message.DisconnectedMessage{message.DISCONNECTED_MESSAGE_TYPE, processor.username}
		toAllUsers(processor, messageToUsers.GetJSON())
	}
}

func identifyCase(processor *ServerProcessor, userName string) {
	userNameAvailable := verifyUserName(userName)
	if userNameAvailable {
		addUser(userName, processor)
		processor.setUserName(userName)
		succes := message.SuccesMessage{message.INFO_MESSAGE_TYPE, "Succes: username has been saved", message.IDENTIFY_MESSAGE_TYPE}
		processor.sendMessage(succes.GetJSON())
	} else {
		str := fmt.Sprintf("username '%s' already used.", userName)
		warning := message.WarningMessageUsername{message.WARNING_MESSAGE_TYPE, str , message.IDENTIFY_MESSAGE_TYPE, userName}
		processor.sendMessage(warning.GetJSON())
	}
}

func userListCase(processor *ServerProcessor) {
	list := getUserList()
	processor.sendMessage(list)
}

func disconnectClientCase(processor *ServerProcessor) {
	processor.disconnectClient()
	disconn := message.DisconnectedMessage{message.DISCONNECTED_MESSAGE_TYPE, processor.username}
	toAllUsers(processor, disconn.GetJSON())
}

func publicMessageCase(processor *ServerProcessor, public string) {
	publicMessage := message.NewMessage{message.PUBLIC_MESSAGE_FROM_TYPE, processor.username, public}
	toAllUsers(processor, publicMessage.GetJSON())	
}

func privateMessageCase(processor *ServerProcessor, receptor string, privateMessage string) {
	err := sendPrivateMessage(receptor, privateMessage, processor.username)
	if err != nil {
		str := fmt.Sprintf("the user '%v' does not exist", receptor)
		warning := message.WarningMessageUsername{message.WARNING_MESSAGE_TYPE, str, message.MESSAGE_TYPE, receptor}
		processor.sendMessage(warning.GetJSON())
	}
}
