package main

import(
	"net"
	"fmt"
	"errors"
	"strings"	
)

type server struct {
	rooms map[string]*room
	users map[string]*ServerProcessor
}

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "1234"
	SERVER_TYPE = "tcp"
)

func (server *server) waitForConnections() {
	fmt.Println("Server is already")
	connectionListener, err := net.Listen(SERVER_TYPE, SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for connections...")
	if err != nil {
		fmt.Println("Something went wrong. :(")
	}
	defer connectionListener.Close()
	
	for {
		connection, err := connectionListener.Accept()
		if err != nil {
			fmt.Println("Connection was not accepted.")
			continue
		}
		fmt.Println("Client connected from", connection.RemoteAddr())
		//ACTIVE status is added by default
		serverProcessor := getServerProcessorInstance(server, connection)
		go serverProcessor.readMessages()
	}
}

func (server *server) initServer() {
	server.rooms = make(map[string]*room)
	server.users = make (map[string]*ServerProcessor)
}

func (server *server) addUser(userName string, processor *ServerProcessor) {
	server.users[userName] = processor
}

func (server *server) addRoom(roomname string, room *room) {
	server.rooms[roomname] = room
}

//sends a message just to users that have been added.
func (server *server) sendMessageToAllUsers(processor *ServerProcessor, message []byte) {
	for _, element := range server.users {
		if processor == element {
			continue
		}
		element.sendMessage(message)
	}	
}

//Returns the identified user list
func (server *server) getUserList() []string {
	var listOfUsers []string
	for username, _ := range server.users{
		listOfUsers = append(listOfUsers, username)
	}
	return listOfUsers
}

//gets the user received
func (server *server) getUserProcessor(userName string)(*ServerProcessor, error){
	if userProcessor, ok := server.users[userName]; ok {
		return userProcessor, nil
	}
	
	errorMessage := fmt.Sprintf("El usuario no existe '%s' ",
		userName)
	return nil, createError(errorMessage)
}


func (server *server) getRoom(roomName string) (*room, error) {
	room, itExists := server.rooms[roomName]
	if !itExists {
		errorString := fmt.Sprintf("El cuarto '%s' no existe",
			roomName)
		return nil, errors.New(errorString)
	}
	return room, nil
}

func (server *server) verifyUserName(userName string) bool {
	if _, ok := server.users[userName]; ok {
		return false
	}
	return true
}

//verifies if the status sent by client is valid.
func (server *server) verifyStatus(status string) (bool) {
	switch status {
	case "AWAY": return true
	case "BUSY": return true
	case "ACTIVE": return true
	default: return false
	}
}

//verifies if the room name is available or not.
func (server *server) verifyRoomName(roomName string) (bool) {
	if _, ok := server.rooms[roomName]; ok {
		return false;
	}
	return true;
}

func (server *server) verifyIdentifiedUsers(users []string) (bool, string) {
	for i := 0; i < len(users); i++ {
		if _, itExists := server.users[users[i]]; !itExists {
			return false, users[i]
		}
	}
	return true, ""
}

func (server *server) deleteUserName(oldName string) {
	delete(server.users, oldName)
}

func (server *server) deleteRoom(roomName string) {
	delete(server.rooms, roomName)
}

func (server *server) createNewRoom(host string, hostProcessor *ServerProcessor, roomName string) (error) {
	if value := strings.Compare(roomName, ""); value == 0 {
		return createError("Nombre de cuarto invalido")
	}
	if isRoomNameValid := server.verifyRoomName(roomName); isRoomNameValid {
		newRoom := getRoomInstance(roomName)
		newRoom.addUser(host)
		server.addRoom(roomName, newRoom)
		return nil
	}
	errorString := fmt.Sprintf("El cuarto '%s' ya existe",
		roomName)
	return createError(errorString)
}

func (server *server) verifyRoomInvitation(host string, roomName string, usersToInvite []string) (error) {
	room, err := server.getRoom(roomName)
	
	if err != nil {
		errorString := fmt.Sprintf("El cuarto '%s' no existe", roomName)
		return createError(errorString)
	}
	if isMember := room.verifyRoomMember(host); !isMember {
		errorString := fmt.Sprintf("El usuario no esta en el cuarto '%s'", roomName)
		return createError(errorString)
	}
	
	return nil
}

func (server *server) disconnectUserFromRoom(userName string, roomName string) (error) {
	room, error1 := server.getRoom(roomName)
	if error1 != nil {
		return error1
	}
	isMember := room.verifyRoomMember(userName)
	if !isMember {
		errorString := fmt.Sprintf("El usuario no se ha unido al cuarto '%s'", roomName)
		return createError(errorString)
	}
	room.removeUser(userName)
	return nil
}

func (server *server) isEmptyRoom(roomName string) (bool) {
	room, _ := server.getRoom(roomName)
	return room.isEmpty()
}

func (server *server) addInvitedUserToRoom(roomName string, userName string, userProcessor *ServerProcessor) {
	room, _ := server.getRoom(roomName)
	room.addInvitedUser(userName)
}

func (server *server) addUserToRoom(userName string, roomName string, userProcessor *ServerProcessor) (error) {
	room, error := server.getRoom(roomName)
	if error != nil {
		return error
	}
	if isInvited := room.verifyInvitedUser(userName); !isInvited {
		errorString := fmt.Sprintf("El usuario no ha sido invitado al cuarto '%s'",
			roomName)
		return createError(errorString)
	}

	if isMember := room.verifyRoomMember(userName); isMember {
		errorString := fmt.Sprintf("El usuario ya se unió al cuarto '%v'",
		roomName)
		return createError(errorString)
	}
	room.addUser(userName)
	return nil
}

func (server *server) removeInvitedUserInRoom(userName string, roomName string) {
	room, _ := server.getRoom(roomName)
	room.removeInvitedUser(userName)
}

func (server *server) sendMessageToRoom(transmitter string, roomName string, message []byte) (error) {
	room, error:= server.getRoom(roomName)
	if  error != nil {
		return error
	}

	if isMember := room.verifyRoomMember(transmitter); !isMember {
		errorString := fmt.Sprintf("El usuario no se ha unido al cuarto '%s'",
			roomName)
		return createError(errorString)
	}
	
	for userName, _ := range(room.users.elements) {
		if val := strings.Compare(userName, transmitter); val == 0 {
			continue
		}
		userProcessor, _ := server.getUserProcessor(userName)
		userProcessor.sendMessage(message)
	}
	return nil
}

func (server *server) getRoomUserList(userName string, roomName string) ([]string, error) {
	room, err := server.getRoom(roomName)
	if err != nil {
		errorString := fmt.Sprintf("El cuarto '%s' no existe",
			roomName)
		return nil, createError(errorString)
	}
	if isMember := room.verifyRoomMember(userName); !isMember {
		errorString := fmt.Sprintf("El usuario no se ha unido al cuarto '%s'",
			roomName)
		return nil , createError(errorString)
	}
	return room.getMemberList(), nil
}

func createError(errorMessage string) (error) {
	return errors.New(errorMessage)
}

