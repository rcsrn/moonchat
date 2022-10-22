package server

import (
	"net"
	"fmt"
	"encoding/json"
	"github.com/rcsrn/moonchat/internal/message"
	"strings"
	"log"
)

type ServerProcessor struct {
	server *Server
	connection net.Conn
	username string
	userStatus string
	identified bool
	rooms []string
	verifier *messageVerifier
	creator *messageCreator
}

func GetServerProcessorInstance(server *Server, connection net.Conn) *ServerProcessor {
	serverProcessor := ServerProcessor{server, connection, "", "ACTIVE", false, make([]string, 1024), &messageVerifier{}, &messageCreator{}}
	return &serverProcessor
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
		messageReceived, error2 := processor.UnmarshalJSON(buffer[:numBitsRead])
		
		if error2 != nil {
			errorWhileReading("unmarshaling", processor, error2)
			break
		}

		if !processor.isValidMessage(messageReceived) {
			disconnectClientCase(processor)
		}
		
		if !processor.identified {
			if value := strings.Compare(messageReceived["type"], message.IDENTIFY_TYPE); value != 0 {
				identifyErrorMessage := processor.creator.getIdentifyErrorMessage()
				processor.sendMessage(identifyErrorMessage)
				disconnectClientCase(processor)
				break
			}
		}

		processor.processMessage(messageReceived)
		log.Printf("message received: %v by %v\n", messageReceived, processor.username)
	}
}

func (processor *ServerProcessor) isValidMessage(messageReceived map[string]string) (bool) {
	processor.verifier.setMessage(messageReceived)
	return processor.verifier.validateMessage()
}

func errorWhileReading(errorCase string, processor *ServerProcessor, error error) {
	log.Printf("Error while %s : %v\n",
		errorCase,
		error.Error())
	
	disconnectClientCase(processor)
	
	if processor.identified {
		log.Printf("Client '%s' disconnected",
			processor.username)
		disconnectedMessage := processor.creator.getDisconnectedMessage(processor.username)
		processor.server.sendMessageToAllUsers(processor, disconnectedMessage)
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
func (processor *ServerProcessor) UnmarshalJSON(j []byte) (map[string]string, error) {
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

func (processor *ServerProcessor) addRoom(newRoom string) {
	processor.rooms = append(processor.rooms, newRoom)
}

func (processor *ServerProcessor) changeStatus(newStatus string) (bool) {
	if accepted := processor.server.VerifyStatus(newStatus); !accepted {
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
		break
	case message.LEAVE_ROOM_TYPE:
		leaveMessageCase(processor, processor.username, gottenMessage["roomname"])
		break
	default: disconnectClientCase(processor)
	}
}

func statusCase(processor *ServerProcessor, newStatus string) {
	if processor.userStatus == newStatus {
		warningString := fmt.Sprintf("El estado ya es '%s'", processor.userStatus)
		warningMessage := processor.creator.getStatusWarningMessage(warningString, message.STATUS_TYPE, processor.userStatus)
		processor.sendMessage(warningMessage)
		return
	}
	
	changedStatus := processor.changeStatus(newStatus)
	if changedStatus {
		succesMessage := processor.creator.getSuccesMessage("succes", message.STATUS_TYPE)
		processor.sendMessage(succesMessage)
		messageToUsers := processor.creator.getNewStatusMessage(processor.username, newStatus)
		processor.server.sendMessageToAllUsers(processor, messageToUsers)
	} else {
		errorMessage := processor.creator.getStatusErrorMessage("estatus invalido", message.STATUS_TYPE)
		processor.sendMessage(errorMessage)
		disconnectClientCase(processor)
		messageToUsers := processor.creator.getDisconnectedMessage(processor.username)
		processor.server.sendMessageToAllUsers(processor, messageToUsers)
	}
}

func identifyCase(processor *ServerProcessor, userName string) {
	if processor.username != "" {
		return
	}
	
	isAvailable := processor.server.VerifyUserName(userName)

	if isAvailable {
		processor.server.AddUser(userName, processor)
		newUserMessage := message.NewUserMessage{message.NEW_USER_TYPE, userName}
		processor.server.sendMessageToAllUsers(processor, newUserMessage.GetJSON())
		
		processor.setUserName(userName)
		
		succesMessage := processor.creator.getSuccesMessage("succes", message.IDENTIFY_TYPE)
		processor.sendMessage(succesMessage)
		
	} else {
		warningString := fmt.Sprintf("El usuario '%s' ya existe", userName)
		warningMessage := processor.creator.getUsernameWarningMessage(warningString, message.IDENTIFY_TYPE, userName)
		processor.sendMessage(warningMessage)
	}
}

func userListCase(processor *ServerProcessor) {
	userList := processor.server.getUserList()
	userListMessage := processor.creator.getUserListMessage(message.USER_LIST_TYPE, userList)
	processor.sendMessage(userListMessage)
}

func disconnectClientCase(processor *ServerProcessor) {
	processor.connection.Close()
	processor.server.deleteUserName(processor.username)
	disconnectedMessage := processor.creator.getDisconnectedMessage(processor.username)
	processor.server.sendMessageToAllUsers(processor, disconnectedMessage)
	processor.leaveAllRooms(processor.username, processor.rooms)
}

func publicMessageCase(processor *ServerProcessor, publicMessageToSend string) {
	publicMessage := processor.creator.getPublicMessage(processor.username, publicMessageToSend)
	processor.server.sendMessageToAllUsers(processor, publicMessage)	
}

func privateMessageCase(processor *ServerProcessor, receptor string, privateMessageToSend string) {
	receptorProcessor, error := processor.server.GetUserProcessor(receptor)
	if error != nil {
		warningMessage := processor.creator.getUsernameWarningMessage(error.Error(), message.MESSAGE_TYPE, receptor)
		processor.sendMessage(warningMessage)
		return
	}
	privateMessage := processor.creator.getPrivateMessage(privateMessageToSend, processor.username)
	receptorProcessor.sendMessage(privateMessage)
}

func newRoomCase(processor *ServerProcessor, roomName string) {
	error := processor.server.CreateNewRoom(processor.username, processor, roomName)
	if error != nil {
		warningMessage := processor.creator.getRoomWarningMessage(error.Error(), message.NEW_ROOM_TYPE, roomName)
		processor.sendMessage(warningMessage)
	} else {
		processor.server.addInvitedUserToRoom(roomName, processor.username, processor)
		succesMessage := processor.creator.getRoomSuccesMessage("succes", message.NEW_ROOM_TYPE, roomName)
		processor.sendMessage(succesMessage)
		processor.addRoom(roomName)
	}
}

func inviteToRoomCase(processor *ServerProcessor, roomName string, users string) {
	usersToInvite := toArrayOfUsers(users)
	if theyAllExist, user := processor.server.VerifyIdentifiedUsers(usersToInvite); !theyAllExist {
		warningString := fmt.Sprintf("El usuario '%s' no existe", user)
		warningMessage := processor.creator.getUsernameWarningMessage(warningString, message.INVITE_TYPE, user)
		processor.sendMessage(warningMessage)
		return
	}
	error:= processor.server.verifyRoomInvitation(processor.username, roomName, usersToInvite)
	if error != nil {
		warningMessage := processor.creator.getRoomWarningMessage(error.Error(), message.INVITE_TYPE, roomName)
		processor.sendMessage(warningMessage)
		return
	}
	processor.sendInvitation(processor.username, roomName, users)
	succesMessage := processor.creator.getRoomSuccesMessage("succes", message.INVITE_TYPE, roomName)
	processor.sendMessage(succesMessage)
}

func (processor *ServerProcessor) sendInvitation(host string, roomName string, users string) {
	ArrayOfUsers := toArrayOfUsers(users)
	for i := 0; i < len(ArrayOfUsers); i++ {
		userProcessor, _ := processor.server.GetUserProcessor(ArrayOfUsers[i])
		processor.server.addInvitedUserToRoom(roomName, userProcessor.username, userProcessor)
		invitationString := fmt.Sprintf("%v te invita al cuarto '%v'",
		host, roomName)
		invitationMessage := processor.creator.getRoomInvitationMessage(invitationString, host, roomName)
		userProcessor.sendMessage(invitationMessage)
	}
}

func joinRoomCase(processor *ServerProcessor, roomName string) {
	error := processor.server.addUserToRoom(processor.username, roomName)
	if error != nil {
		warningMessage := processor.creator.getRoomWarningMessage(error.Error(), message.JOIN_ROOM_TYPE, roomName)
		processor.sendMessage(warningMessage)
		return
	}
	
	processor.server.removeInvitedUserInRoom(processor.username, roomName)
	joinedMessage := processor.creator.getJoinedMessage(roomName, processor.username)
	processor.server.sendMessageToRoom(processor.username, roomName, joinedMessage)
	succesMessage := processor.creator.getRoomSuccesMessage("succes", message.JOIN_ROOM_TYPE, roomName)
	processor.sendMessage(succesMessage)
	processor.addRoom(roomName)
}

func roomUsersCase(processor *ServerProcessor, roomName string) {
	roomUserList, err := processor.server.getRoomUserList(processor.username, roomName)
	if err != nil {
		warningMessage := processor.creator.getRoomWarningMessage(err.Error(), message.ROOM_USERS_TYPE, roomName)
		processor.sendMessage(warningMessage)
		return
	}
	roomUsersMessage := processor.creator.getUserListMessage(message.ROOM_USER_LIST_TYPE, roomUserList)
	processor.sendMessage(roomUsersMessage)
}

//auxiliar function to convert this line to an array of users. 
func toArrayOfUsers(line string) ([]string) {
	lines := strings.Split(line, ",")
	return lines
}

func roomMessageCase(processor *ServerProcessor, roomName string, messageToSend string) {
	messageToSend = fmt.Sprintf("[%v]: ", roomName) + messageToSend
	
	roomMessage := processor.creator.getRoomMessage(roomName, processor.username, messageToSend)
	
	error := processor.server.sendMessageToRoom(processor.username, roomName, roomMessage)
	if error != nil {
		warningMessage := processor.creator.getRoomWarningMessage(error.Error(), message.ROOM_MESSAGE_TYPE ,roomName)
		processor.sendMessage(warningMessage)
		return
	}
}

func leaveMessageCase(processor *ServerProcessor, userName string, roomName string) {
	error := processor.server.disconnectUserFromRoom(userName, roomName)
	if error != nil {
		warningMessage := processor.creator.getRoomWarningMessage(error.Error(), message.LEAVE_ROOM_TYPE, roomName)
		processor.sendMessage(warningMessage)
		return
	}

	if isEmpty := processor.server.isEmptyRoom(roomName); isEmpty {
		processor.server.deleteRoom(roomName)
	}

	succesMessage := processor.creator.getRoomSuccesMessage("succes", message.LEAVE_ROOM_TYPE, roomName)
	processor.sendMessage(succesMessage)

	leftRoomMessage := processor.creator.getLeftRoomMessage(roomName, userName)
	processor.server.sendMessageToRoom(userName, roomName, leftRoomMessage)
	
}

func (processor *ServerProcessor) leaveAllRooms(userName string, rooms []string) {
	for i := 0; i < len(rooms); i++ {
		processor.server.disconnectUserFromRoom(userName, rooms[i])
		leftRoomMessage := processor.creator.getLeftRoomMessage(rooms[i], userName)
		processor.server.sendMessageToRoom(userName, rooms[i], leftRoomMessage)
	}
}
