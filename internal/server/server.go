package server

import(
	"net"
	"fmt"
	"errors"
)

type Server struct {
	rooms map[string]*Room
	users map[string]*ServerProcessor
}

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "1234"
	SERVER_TYPE = "tcp"
)

func GetServerInstance() *Server {
	return &Server{make(map[string]*Room), make (map[string]*ServerProcessor)}
}

func (server *Server) GetRooms() map[string]*Room {
	return server.rooms
}

func (server *Server) GetUsers() map[string]*ServerProcessor {
	return server.users
}

func (server *Server) waitForConnections() {
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
		serverProcessor := GetServerProcessorInstance(server, connection)
		go serverProcessor.readMessages()
	}
}

func (server *Server) AddUser(userName string, processor *ServerProcessor) {
	server.users[userName] = processor
}

func (server *Server) AddRoom(roomname string, room *Room) {
	server.rooms[roomname] = room
}

//sends a message just to users that have been added.
func (server *Server) sendMessageToAllUsers(processor *ServerProcessor, message []byte) {
	for _, element := range server.users {
		if processor == element {
			continue
		}
		element.sendMessage(message)
	}	
}

//Returns the identified user list
func (server *Server) getUserList() []string {
	var listOfUsers []string
	for username, _ := range server.users{
		listOfUsers = append(listOfUsers, username)
	}
	return listOfUsers
}

//gets the user received
func (server *Server) GetUserProcessor(userName string)(*ServerProcessor, error){
	if userProcessor, ok := server.users[userName]; ok {
		return userProcessor, nil
	}
	
	errorMessage := fmt.Sprintf("El usuario no existe '%s' ",
		userName)
	return nil, createError(errorMessage)
}


func (server *Server) GetRoom(roomName string) (*Room, error) {
	room, itExists := server.rooms[roomName]
	if !itExists {
		errorString := fmt.Sprintf("El cuarto '%s' no existe",
			roomName)
		return nil, errors.New(errorString)
	}
	return room, nil
}

func (server *Server) VerifyUserName(userName string) bool {
	if _, ok := server.users[userName]; ok {
		return false
	}
	return true
}

//verifies if the status sent by client is valid.
func (server *Server) VerifyStatus(status string) (bool) {
	switch status {
	case "AWAY": return true
	case "BUSY": return true
	case "ACTIVE": return true
	default: return false
	}
}

//verifies if the room name is available or not.
func (server *Server) VerifyRoomName(roomName string) (bool) {
	if _, ok := server.rooms[roomName]; ok {
		return false;
	}
	return true;
}

func (server *Server) VerifyIdentifiedUsers(users []string) (bool, string) {
	for i := 0; i < len(users); i++ {
		if _, itExists := server.users[users[i]]; !itExists {
			return false, users[i]
		}
	}
	return true, ""
}

func (server *Server) deleteUserName(oldName string) {
	delete(server.users, oldName)
}

func (server *Server) deleteRoom(roomName string) {
	delete(server.rooms, roomName)
}

func (server *Server) CreateNewRoom(host string, hostProcessor *ServerProcessor, roomName string) (error) {
	if roomName == "" {
		return createError("Nombre de cuarto invalido")
	}
	if isAvailable := server.VerifyRoomName(roomName); isAvailable {
		newRoom := GetRoomInstance(roomName)
		newRoom.AddUser(host)
		server.AddRoom(roomName, newRoom)
		return nil
	}
	errorString := fmt.Sprintf("El cuarto '%s' ya existe",
		roomName)
	return createError(errorString)
}

func (server *Server) verifyRoomInvitation(host string, roomName string, usersToInvite []string) (error) {
	room, err := server.GetRoom(roomName)
	
	if err != nil {
		errorString := fmt.Sprintf("El cuarto '%s' no existe", roomName)
		return createError(errorString)
	}
	if isMember := room.VerifyRoomMember(host); !isMember {
		errorString := fmt.Sprintf("El usuario no esta en el cuarto '%s'", roomName)
		return createError(errorString)
	}
	
	return nil
}

func (server *Server) disconnectUserFromRoom(userName string, roomName string) (error) {
	room, error1 := server.GetRoom(roomName)
	if error1 != nil {
		return error1
	}
	isMember := room.VerifyRoomMember(userName)
	if !isMember {
		errorString := fmt.Sprintf("El usuario no se ha unido al cuarto '%s'", roomName)
		return createError(errorString)
	}
	room.RemoveUser(userName)
	return nil
}

func (server *Server) isEmptyRoom(roomName string) (bool) {
	room, _ := server.GetRoom(roomName)
	return room.IsEmpty()
}

func (server *Server) addInvitedUserToRoom(roomName string, userName string, userProcessor *ServerProcessor) {
	room, _ := server.GetRoom(roomName)
	room.AddInvitedUser(userName)
}

func (server *Server) addUserToRoom(userName string, roomName string) (error) {
	room, error := server.GetRoom(roomName)
	if error != nil {
		return error
	}
	if isInvited := room.VerifyInvitedUser(userName); !isInvited {
		errorString := fmt.Sprintf("El usuario no ha sido invitado al cuarto '%s'",
			roomName)
		return createError(errorString)
	}

	if isMember := room.VerifyRoomMember(userName); isMember {
		errorString := fmt.Sprintf("El usuario ya se unió al cuarto '%v'",
		roomName)
		return createError(errorString)
	}
	room.AddUser(userName)
	return nil
}

func (server *Server) removeInvitedUserInRoom(userName string, roomName string) {
	room, _ := server.GetRoom(roomName)
	room.RemoveInvitedUser(userName)
}

func (server *Server) sendMessageToRoom(transmitter string, roomName string, message []byte) (error) {
	room, error:= server.GetRoom(roomName)
	if  error != nil {
		return error
	}

	if isMember := room.VerifyRoomMember(transmitter); !isMember {
		errorString := fmt.Sprintf("El usuario no se ha unido al cuarto '%s'",
			roomName)
		return createError(errorString)
	}
	
	for userName, _ := range(room.GetMembers()) {
		if userName == transmitter {
			continue
		}
		userProcessor, _ := server.GetUserProcessor(userName)
		userProcessor.sendMessage(message)
	}
	return nil
}

func (server *Server) getRoomUserList(userName string, roomName string) ([]string, error) {
	room, err := server.GetRoom(roomName)
	if err != nil {
		errorString := fmt.Sprintf("El cuarto '%s' no existe",
			roomName)
		return nil, createError(errorString)
	}
	if isMember := room.VerifyRoomMember(userName); !isMember {
		errorString := fmt.Sprintf("El usuario no se ha unido al cuarto '%s'",
			roomName)
		return nil , createError(errorString)
	}
	return room.GetMemberList(), nil
}

func createError(errorMessage string) (error) {
	return errors.New(errorMessage)
}

