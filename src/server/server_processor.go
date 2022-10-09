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
				disconnectedMessage := getDisconnectedMessage(processor.username)
				toAllUsers(processor, disconnectedMessage)
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
		warningString := fmt.Sprintf("The current status is '%s'. Please select other.", processor.userStatus)
		warningMessage := getStatusWarningMessage(warningString, message.STATUS_TYPE, processor.userStatus)
		processor.sendMessage(warningMessage)
		return
	}
	
	changedStatus := processor.changeStatus(newStatus)
	if changedStatus {
		succesMessage := getSuccesMessage("Succes: status changed succesfully", message.STATUS_TYPE)
		processor.sendMessage(succesMessage)
		messageToUsers := getNewStatusMessage(processor.username, newStatus)
		toAllUsers(processor, messageToUsers)
	} else {
		errorMessage := getStatusErrorMessage("Invalid status!", message.STATUS_TYPE)
		processor.sendMessage(errorMessage)
		processor.disconnectClient()
		messageToUsers := getDisconnectedMessage(processor.username)
		toAllUsers(processor, messageToUsers)
	}
}

func identifyCase(processor *ServerProcessor, userName string) {
	TheUserNameIsAvailable := verifyUserName(userName)

	if TheUserNameIsAvailable {
		if itHasOldName := strings.Compare(processor.username, ""); itHasOldName != 0 {
			removeOldName(processor.username)
		}
		addUser(userName, processor)
		processor.setUserName(userName)
		succesMessage := getSuccesMessage("Succes: username has been saved", message.IDENTIFY_TYPE)
		processor.sendMessage(succesMessage)
	} else {
		warningString := fmt.Sprintf("username '%s' is already used.", userName)
		warningMessage := getUsernameWarningMessage(warningString, message.IDENTIFY_TYPE, userName)
		processor.sendMessage(warningMessage)
	}
}

func userListCase(processor *ServerProcessor) {
	userList := getUserList()
	userListMessage := getUserListMessage(message.USER_LIST_TYPE, userList)
	processor.sendMessage(userListMessage)
}

func disconnectClientCase(processor *ServerProcessor) {
	processor.disconnectClient()
	disconnectedMessage := getDisconnectedMessage(processor.username)
	toAllUsers(processor, disconnectedMessage)
}

func publicMessageCase(processor *ServerProcessor, publicMessageToSend string) {
	publicMessage := getPublicMessage(processor.username, publicMessageToSend)
	toAllUsers(processor, publicMessage)	
}

func privateMessageCase(processor *ServerProcessor, receptor string, privateMessageToSend string) {
	receptorProcessor, error := getUserProcessor(receptor)
	if error != nil {
		warningMessage := getUsernameWarningMessage(error.Error(), message.MESSAGE_TYPE, receptor)
		processor.sendMessage(warningMessage)
		return
	}
	privateMessage := getPrivateMessage(privateMessageToSend, processor.username)
	receptorProcessor.sendMessage(privateMessage)
}

func newRoomCase(processor *ServerProcessor, roomName string) {
	message, _ := createNewRoom(processor.username, processor, roomName)
	processor.sendMessage(message)
}

func inviteToRoomCase(processor *ServerProcessor, roomName string, users string) {
	usersToInvite := toArrayOfUsers(users)
	if theyAllExist, user := verifyIdentifiedUsers(usersToInvite); !theyAllExist {
		warningString := fmt.Sprintf("The user '%s' does not exist", user)
		warningMessage := getUsernameWarningMessage(warningString, message.INVITE_TYPE, user)
		processor.sendMessage(warningMessage)
		return
	}
	error:= inviteUsersToRoom(processor.username, roomName, usersToInvite)
	if error != nil {
		warningMessage := getRoomWarningMessage(error.Error(), message.INVITE_TYPE, roomName)
		processor.sendMessage(warningMessage)
		return
	}
	succesMessage := getRoomSuccesMessage("Succes: users have been invited to room", message.INVITE_TYPE, roomName)
	processor.sendMessage(succesMessage)
}

func joinRoomCase(processor *ServerProcessor, roomName string) {
	err := joinRoom(processor.username, roomName)
	if err != nil {
		processor.sendMessage(getRoomWarningMessage(err.Error(), message.JOIN_ROOM_TYPE, roomName))
		return
	}
	succesString := fmt.Sprintf("Succes: you have been added to room '%s'!",
		roomName)
	succesMessage := getRoomSuccesMessage(succesString, message.JOIN_ROOM_TYPE, roomName)
	processor.sendMessage(succesMessage)
}

func roomUsersCase(processor *ServerProcessor, roomName string) {
	roomUserList, err := getRoomUserList(processor.username, roomName)
	if err != nil {
		warningMessage := getRoomWarningMessage(err.Error(), message.ROOM_USERS_TYPE, roomName)
		processor.sendMessage(warningMessage)
		return
	}
	roomUsersMessage := getUserListMessage(message.ROOM_USER_LIST_TYPE, roomUserList)
	processor.sendMessage(roomUsersMessage)
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

func getSuccesMessage(succes string, operation string) ([]byte) {
	succesMessage := message.SuccesMessage{message.INFO_TYPE, succes, operation}
	return succesMessage.GetJSON()
}

func getUsernameWarningMessage(warning string, operation string, userName string) ([]byte) {
	warningMessage := message.UsernameWarningMessage{message.WARNING_TYPE, warning, operation, userName}
	return warningMessage.GetJSON()
}

func getStatusWarningMessage(warning string, operation string, status string) ([]byte) {
	warningMessage := message.StatusWarningMessage{message.WARNING_TYPE, warning, operation, status}
	return warningMessage.GetJSON()
}

func getRoomWarningMessage(warning string, operation string, roomName string) ([]byte) {
	warningMessage := message.RoomWarningMessage{message.WARNING_TYPE, warning, operation, roomName}
	return warningMessage.GetJSON()
}

func getDisconnectedMessage(userName string) ([]byte) {
	disconnectedMessage := message.DisconnectedMessage{message.DISCONNECTED_TYPE, userName}
	return disconnectedMessage.GetJSON()
}

func getUserListMessage(typeOfList string, userList []string) ([]byte) {
	switch typeOfList {
	case message.ROOM_USER_LIST_TYPE:
		userListMessage := message.UserList{message.ROOM_USER_LIST_TYPE, userList}
		return userListMessage.GetJSON()
	default:
		userListMessage := message.UserList{message.USER_LIST_TYPE, userList}
		return userListMessage.GetJSON()
	}
}

func getPrivateMessage(private string, transmitter string) ([]byte) {
	privateMessage := message.NewMessage{message.PRIVATE_TYPE, transmitter, private}
	return privateMessage.GetJSON()
}

func getPublicMessage(userName string, public string) ([]byte) {
	publicMessage := message.NewMessage{message.PUBLIC_MESSAGE_FROM_TYPE, userName, public}
	return publicMessage.GetJSON()
}

func getNewStatusMessage(userName string, status string) ([]byte) {
	newStatusMessage := message.NewStatusMessage{message.NEW_STATUS_TYPE, userName, status}
	return newStatusMessage.GetJSON()
}

func getStatusErrorMessage(error string, operation string) ([]byte) {
	errorMessage := message.StatusErrorMessage{message.ERROR_TYPE, error, operation}
	return errorMessage.GetJSON()
}
