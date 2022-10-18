package main

import (
	"net"
	"fmt"
	"encoding/json"
	"github.com/rcsrn/moonchat/cmd/message"
	"strings"
	"log"
)

type ServerProcessor struct {
	server *server
	connection net.Conn
	username string
	userStatus string
	identified bool
	rooms []string
}

func getServerProcessorInstance(server *server, connection net.Conn) *ServerProcessor {
	server := ServerProcessor{server,
		connection, "", "ACTIVE", false, make([]string, 512)}
	return &server
}

// reads sent messages by client
func (processor *ServerProcessor) readMessages() {
	for {
		buffer := make([]byte, 1024)
		numBitsRead, error1 := processor.connection.Read(buffer)
		
		if error1 != nil {
			errorWhileReading("reading", processor, error1)
			break
		}

		//It is necessary to specify the exact number of bits read by Read() method to unmarshal.
		messageReceived, error2 := processor.unmarshalJSON(buffer[:numBitsRead])
		
		if error2 != nil {
			errorWhileReading("unmarshaling", processor, error2)
			break
		}

		if !processor.identified {
			if value := strings.Compare(messageReceived["type"], message.IDENTIFY_TYPE); value != 0 {
				processor.sendMessage(getIdentifyErrorMessage())
				disconnectClientCase(processor)
				break
			}
		}
		
		processor.processMessage(messageReceived)
		log.Printf("message received: %v by %v\n", messageReceived, processor.username)
	}
}

func errorWhileReading(errorCase string, processor *ServerProcessor, error error) {
	log.Printf("Error while %s : %v\n",
		errorCase,
		error.Error())
	
	disconnectClientCase(processor)
	
	if processor.identified {
		log.Printf("Client '%s' disconnected",
			processor.username)
		disconnectedMessage := getDisconnectedMessage(processor.username)
		sendMessageToAllUsers(processor, disconnectedMessage)
	} else {
		log.Printf("Client disconnected")
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

//sends messages to the client through its connection.
func (processor *ServerProcessor) sendMessage(message []byte) {
	if processor.connection != nil {
		processor.connection.Write(message)
	}
}

//unmarshals messages that need to be processed by serverProcessor
func (processor *ServerProcessor) unmarshalJSON(j []byte) (map[string]string, error) {
	var mapString map[string]string 
	error1 := json.Unmarshal(j, &mapString)
	if error1 != nil {
		if val := strings.Compare(error1.Error(), "json: cannot unmarshal array into Go value of type string"); val == 0 {
			var message message.InviteToRoomMessage
			error2 := json.Unmarshal(j, &message)
			if error2 != nil {
				return nil, error2
			}
			return convertMessageToMapString(message), nil
		} else {
			return nil, error1
		}
	}
	return mapString, nil
}

//auxiliar function to deal the case when it is necessary to work with an array of strings.
func convertMessageToMapString(message message.InviteToRoomMessage) (map[string]string) {
	mapString := make(map[string]string)
	mapString["type"] = message.Type
	mapString["roomname"] = message.Roomname
	mapString["usernames"] = strings.Join(message.Usernames, ",")
	return mapString
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
	case message.ROOM_MESSAGE_TYPE:
		roomMessageCase(processor, gottenMessage["roomname"], gottenMessage["message"])
	case message.LEAVE_ROOM_TYPE:
		leaveMessageCase(processor, processor.username, gottenMessage["roomname"])
	default: disconnectClientCase(processor)
	}
}

func statusCase(processor *ServerProcessor, newStatus string) {
	if value := strings.Compare(processor.userStatus, newStatus); value == 0 {
		warningString := fmt.Sprintf("El estado ya es '%s'", processor.userStatus)
		warningMessage := getStatusWarningMessage(warningString, message.STATUS_TYPE, processor.userStatus)
		processor.sendMessage(warningMessage)
		return
	}
	
	changedStatus := processor.changeStatus(newStatus)
	if changedStatus {
		succesMessage := getSuccesMessage("succes", message.STATUS_TYPE)
		processor.sendMessage(succesMessage)
		messageToUsers := getNewStatusMessage(processor.username, newStatus)
		sendMessageToAllUsers(processor, messageToUsers)
	} else {
		errorMessage := getStatusErrorMessage("estatus invalido", message.STATUS_TYPE)
		processor.sendMessage(errorMessage)
		disconnectClientCase(processor)
		messageToUsers := getDisconnectedMessage(processor.username)
		sendMessageToAllUsers(processor, messageToUsers)
	}
}

func identifyCase(processor *ServerProcessor, userName string) {
	isAvailable := verifyUserName(userName)

	if isAvailable {
		if itHasOldName := strings.Compare(processor.username, ""); itHasOldName != 0 {
			deleteUserName(processor.username)
		}
		addUser(userName, processor)
		processor.setUserName(userName)
		succesMessage := getSuccesMessage("succes", message.IDENTIFY_TYPE)
		processor.sendMessage(succesMessage)
	} else {
		warningString := fmt.Sprintf("El usuario '%s' ya existe", userName)
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
	processor.connection.Close()
	deleteUserName(processor.username)
	disconnectedMessage := getDisconnectedMessage(processor.username)
	sendMessageToAllUsers(processor, disconnectedMessage)
	leaveAllRooms(processor.username, processor.rooms)
}

func publicMessageCase(processor *ServerProcessor, publicMessageToSend string) {
	publicMessage := getPublicMessage(processor.username, publicMessageToSend)
	sendMessageToAllUsers(processor, publicMessage)	
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
	error := createNewRoom(processor.username, processor, roomName)
	if error != nil {
		warningMessage := getRoomWarningMessage(error.Error(), message.NEW_ROOM_TYPE, roomName)
		processor.sendMessage(warningMessage)
	} else {
		addInvitedUserToRoom(roomName, processor.username, processor)
		succesMessage := getRoomSuccesMessage("succes", message.NEW_ROOM_TYPE, roomName)
		processor.sendMessage(succesMessage)
	}
}

func inviteToRoomCase(processor *ServerProcessor, roomName string, users string) {
	usersToInvite := toArrayOfUsers(users)
	if theyAllExist, user := verifyIdentifiedUsers(usersToInvite); !theyAllExist {
		warningString := fmt.Sprintf("El usuario '%s' no existe", user)
		warningMessage := getUsernameWarningMessage(warningString, message.INVITE_TYPE, user)
		processor.sendMessage(warningMessage)
		return
	}
	error:= verifyRoomInvitation(processor.username, roomName, usersToInvite)
	if error != nil {
		warningMessage := getRoomWarningMessage(error.Error(), message.INVITE_TYPE, roomName)
		processor.sendMessage(warningMessage)
		return
	}
	sendInvitation(processor.username, roomName, users)
	succesMessage := getRoomSuccesMessage("succes", message.INVITE_TYPE, roomName)
	processor.sendMessage(succesMessage)
}

func sendInvitation(host string, roomName string, users string) {
	ArrayOfUsers := toArrayOfUsers(users)
	for i := 0; i < len(ArrayOfUsers); i++ {
		userProcessor, _ := getUserProcessor(ArrayOfUsers[i])
		addInvitedUserToRoom(roomName, userProcessor.username, userProcessor)
		invitationString := fmt.Sprintf("%v te invita al cuarto '%v'",
		host, roomName)
		invitationMessage := getRoomInvitationMessage(invitationString, host, roomName)
		userProcessor.sendMessage(invitationMessage)
	}
}

func joinRoomCase(processor *ServerProcessor, roomName string) {
	error := addUserToRoom(processor.username, roomName, processor)
	if error != nil {
		processor.sendMessage(getRoomWarningMessage(error.Error(), message.JOIN_ROOM_TYPE, roomName))
		return
	}
	
	removeInvitedUserInRoom(processor.username, roomName)
	joinedMessage := getJoinedMessage(roomName, processor.username)
	sendMessageToRoom(processor.username, roomName, joinedMessage)
	succesString := fmt.Sprintf("succes",
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
	lines := strings.Split(line, ",")
	return lines
}

func roomMessageCase(processor *ServerProcessor, roomName string, messageToSend string) {
	messageToSend = fmt.Sprintf("[%v]: ", roomName) + messageToSend
	
	messageToRoomUsers := getRoomMessage(roomName, processor.username, messageToSend)
	
	error := sendMessageToRoom(processor.username, roomName, messageToRoomUsers)
	if error != nil {
		warningMessage := getRoomWarningMessage(error.Error(), message.ROOM_MESSAGE_TYPE ,roomName)
		processor.sendMessage(warningMessage)
		return
	}
}

func leaveMessageCase(processor *ServerProcessor, userName string, roomName string) {
	error := disconnectUserFromRoom(userName, roomName)
	if error != nil {
		warningMessage := getRoomWarningMessage(error.Error(), message.LEAVE_ROOM_TYPE, roomName)
		processor.sendMessage(warningMessage)
		return
	}

	if isEmpty := isEmptyRoom(roomName); isEmpty {
		deleteRoom(roomName)
	}

	succesMessage := getRoomSuccesMessage("succes", message.LEAVE_ROOM_TYPE, roomName)
	processor.sendMessage(succesMessage)

	leftRoomMessage := getLeftRoomMessage(roomName, userName)
	sendMessageToRoom(userName, roomName, leftRoomMessage)
	
}

func leaveAllRooms(userName string, rooms []string) {
	for i := 0; i < len(rooms); i++ {
		disconnectUserFromRoom(userName, rooms[i])
		leftRoomMessage := getLeftRoomMessage(rooms[i], userName)
		sendMessageToRoom(userName, rooms[i], leftRoomMessage)
	}
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

func getRoomInvitationMessage(invitation string, host string, roomName string) ([]byte) {
	roomInvitationMessage := message.RoomInvitationMessage{message.INVITATION_TYPE, invitation, host, roomName}
	return roomInvitationMessage.GetJSON()
}

func getJoinedMessage(roomName string, userName string) ([]byte) {
	joinedMessage := message.JoinedRoomMessage{message.JOINED_ROOM_TYPE, roomName, userName}
	return joinedMessage.GetJSON()
}

func getRoomMessage(roomName string, userName string, messageToSend string) ([]byte) {
	roomMessage := message.RoomMessage{message.ROOM_MESSAGE_FROM_TYPE, roomName, userName, messageToSend}
	return roomMessage.GetJSON()
}

func getLeftRoomMessage(roomName string, userName string) ([]byte) {
	leftRoomMessage := message.LeftRoomMessage{message.LEFT_ROOM_TYPE, roomName, userName}
	return leftRoomMessage.GetJSON()
}

func getIdentifyErrorMessage() ([]byte) {
	errorMessage := message.IdentifyErrorMessage{message.ERROR_TYPE, "User not identified"}
	return errorMessage.GetJSON()
}
