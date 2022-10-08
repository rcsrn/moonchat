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
	identified bool
}


// reads sent messages by client
func (processor *ServerProcessor) readMessages() {
	for {
		buffer := make([]byte, 1024)
		length, err1 := processor.connection.Read(buffer)
		
		if err1 != nil {
			fmt.Println("Error while reading:", err1.Error())
			processor.sendMessage([]byte("Sorry, something went wrong :(\n"))
			processor.disconnectClient()
			break
		}
		
		messageRecieved, err2 := processor.unmarshalJSON(buffer[:length])
		
		if err2 != nil {
			processor.sendMessage([]byte("Sorry, bad operation...\n"))
			processor.disconnectClient()
			if processor.identified {
				disconn := message.DisconnectedMessage{message.DISCONNECTED_TYPE, processor.username}
				toAllUsers(processor, disconn.GetJSON())
			}
			log.Printf("Error while unmarshaling: %v\n", err2.Error())
			if processor.identified == true {
				log.Printf("Client '%s' disconnected", processor.username)
			} else {
				log.Printf("Client disconnected")
			}
			break
		}
		
		processor.processMessage(messageRecieved)
		fmt.Printf("message received: %v by %v\n", messageRecieved, processor.username)
	}
}

//sets the username to the client once it has been identified.
func (processor *ServerProcessor) setUserName(name string) {
	processor.username = name
	processor.identified = true
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
	case message.IDENTIFY_TYPE:
		identifyCase(processor, gottenMessage["username"])
		break
	case message.STATUS_TYPE:
		statusCase(processor, gottenMessage["status"])
		break
	case message.USERS_TYPE:
		userListCase(processor)
		break
	case message.DISCONNECT_TYPE:
		disconnectClientCase(processor)
		break
	case message.PUBLIC_TYPE:
		publicMessageCase(processor, gottenMessage["message"])
		break
	case message.MESSAGE_TYPE:
		privateMessageCase(processor, gottenMessage["username"], gottenMessage["message"])
		break
	case message.NEW_ROOM_TYPE:
		newRoomCase(processor, gottenMessage["roomname"])
		break
	case message.INVITE_TYPE:
		inviteToRoomCase(processor, gottenMessage["roomname"], gottenMessage["usernames"])
		break
	case message.JOIN_ROOM_TYPE:
		joinRoomCase(processor, gottenMessage["roomname"])
		break
	case message.ROOM_USERS_TYPE:
		roomUsersCase(processor, gottenMessage["roomname"])
		break
	}
}

func statusCase(processor *ServerProcessor, newStatus string) {
	if value := strings.Compare(processor.userStatus, newStatus); value == 0 {
		str := fmt.Sprintf("The current status is '%s'. Please select other.", processor.userStatus)
		warning := message.StatusWarningMessage{message.WARNING_TYPE, str, message.STATUS_TYPE, processor.userStatus}
		processor.sendMessage(warning.GetJSON())
		return
	}
	
	changedStatus := processor.changeStatus(newStatus)
	if changedStatus {
		succes := message.SuccesMessage{message.INFO_TYPE, "Succes: status changed succesfully", message.STATUS_TYPE}
		processor.sendMessage(succes.GetJSON())
		messageToUsers := message.NewStatusMessage{message.NEW_STATUS_TYPE, processor.username, newStatus}
		toAllUsers(processor, messageToUsers.GetJSON())
	} else {
		error := message.ErrorMessageStatus{message.ERROR_TYPE, "Invalid status!", message.STATUS_TYPE}
		processor.sendMessage(error.GetJSON())
		processor.disconnectClient()
		messageToUsers := message.DisconnectedMessage{message.DISCONNECTED_TYPE, processor.username}
		toAllUsers(processor, messageToUsers.GetJSON())
	}
}

func identifyCase(processor *ServerProcessor, userName string) {
	TheUserNameIsAvailable := verifyUserName(userName)

	if TheUserNameIsAvailable {
		addUser(userName, processor)
		processor.setUserName(userName)
		succes := message.SuccesMessage{message.INFO_TYPE, "Succes: username has been saved", message.IDENTIFY_TYPE}
		processor.sendMessage(succes.GetJSON())
	} else {
		str := fmt.Sprintf("username '%s' already used.", userName)
		warning := message.UsernameWarningMessage{message.WARNING_TYPE, str , message.IDENTIFY_TYPE, userName}
		processor.sendMessage(warning.GetJSON())
	}
}

func userListCase(processor *ServerProcessor) {
	list := getUserList()
	processor.sendMessage(list)
}

func disconnectClientCase(processor *ServerProcessor) {
	processor.disconnectClient()
	disconn := message.DisconnectedMessage{message.DISCONNECTED_TYPE, processor.username}
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
		warning := message.UsernameWarningMessage{message.WARNING_TYPE, str, message.MESSAGE_TYPE, receptor}
		processor.sendMessage(warning.GetJSON())
	}
}

func newRoomCase(processor *ServerProcessor, roomName string) {
	message, _ := createNewRoom(processor.username, processor, roomName)
	processor.sendMessage(message)
}

func inviteToRoomCase(processor *ServerProcessor, roomName string, users string) {
	usersToInvite := toArrayOfUsers(users)
	if theyAllExist, user := verifyIdentifiedUsers(usersToInvite); !theyAllExist {
		warningStr := fmt.Sprintf("The user '%s' does not exist", user)
		warningMessage := message.UsernameWarningMessage{message.WARNING_TYPE, warningStr, message.INVITE_TYPE, user}
		processor.sendMessage(warningMessage.GetJSON())
		return
	}
	error:= inviteUsersToRoom(processor.username, roomName, usersToInvite)
	if error != nil {
		warningMessage := message.RoomWarningMessage{message.WARNING_TYPE, error.Error(), message.INVITE_TYPE, roomName}
		processor.sendMessage(warningMessage.GetJSON())
		return
	}
	succesMessage := message.RoomSuccesMessage{message.INFO_TYPE, "Succes: users have been invited to room", message.INVITE_TYPE, roomName}
	processor.sendMessage(succesMessage.GetJSON())
}

func joinRoomCase(processor *ServerProcessor, roomName string) {
	err := joinRoom(processor.username, roomName)
	if err != nil {
		warningMessage := message.RoomWarningMessage{message.WARNING_TYPE, err.Error(), message.JOIN_ROOM_TYPE, roomName}
		processor.sendMessage(warningMessage.GetJSON())
		return
	}
	succesString := fmt.Sprintf("Succes: you have been added to room '%s'!",
		roomName)
	succesMessage := message.RoomSuccesMessage{message.INFO_TYPE, succesString, message.JOIN_ROOM_TYPE, roomName}
	processor.sendMessage(succesMessage.GetJSON())
}

func roomUsersCase(processor *ServerProcessor, roomName string) {
	memberList, err := getRoomUserList(processor.username, roomName)
	if err != nil {
		warningMessage := message.RoomWarningMessage{message.WARNING_TYPE, err.Error(), message.ROOM_USERS_TYPE, roomName}
		processor.sendMessage(warningMessage.GetJSON())
		return
	}
	roomUsersMessage := message.UserList{message.ROOM_USERS_TYPE, memberList}
	processor.sendMessage(roomUsersMessage.GetJSON())
}

//auxiliar function to convert this line to an array of users. 
func toArrayOfUsers(line string) ([]string) {
	line = line[1:len(line) - 1]
	lines := strings.Split(line, ",")
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.Trim(lines[i], " ")
		lines[i] = lines[i][1:len(lines[i])]
	}
	return lines
}

func getRoomSuccesMessage(succes string, operation string, roomName string) ([]byte) {
	succesMessage := message.RoomSuccesMessage{message.INFO_TYPE, succes, operation, roomName}
	return succesMessage.GetJSON()
}

func getSuccesMessage(succes, operation string) ([]byte) {
	succesMessage := message.SuccesMessage{message.INFO_TYPE, succes, operation}
	return succesMessage.GetJSON()
}

func UsernameWarningMessage(warning string, operation string, userName string) ([]byte) {
	warningMessage := message.UsernameWarningMessage{message.WARNING_TYPE, warning, operation, userName}
	return warningMessage.GetJSON()
}

func StatusWarningMessage(warning string, operation string, status string) ([]byte) {
	warningMessage := message.StatusWarningMessage{message.WARNING_TYPE, warning, operation, status}
	return warningMessage.GetJSON()
}

func RoomWarningMessage(warning string, operation string, roomName string) ([]byte) {
	warningMessage := message.RoomWarningMessage{message.WARNING_TYPE, warning, operation, roomName}
	return warningMessage.GetJSON()
}
