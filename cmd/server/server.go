package main

import(
	"net"
	"fmt"
	"github.com/rcsrn/moonchat/cmd/message"
	"sync"
	"errors"
	"strings"	
)

type server struct {
}

//Maps protected for concurrency
var counter = struct{
	blocker sync.RWMutex
	rooms map[string]*room
	users map[string]*ServerProcessor
}{users: make(map[string]*ServerProcessor)}

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
		serverProcessor := getServerProcessorInstance(connection)
		go serverProcessor.readMessages()
	}
}

func (server *server) initRooms() {
	counter.rooms = make(map[string]*room)
}

func addUser(userName string, processor *ServerProcessor) {
	counter.users[userName] = processor
	m := message.NewUserMessage{message.NEW_USER_TYPE, userName}
	sendMessageToAllUsers(processor, m.GetJSON())
}

func addRoom(roomname string, room *room) {
	counter.rooms[roomname] = room
}

//sends a message just to users that have been added.
func sendMessageToAllUsers(processor *ServerProcessor, message []byte) {
	for _, element := range counter.users {
		if processor == element {
			continue
		}
		element.sendMessage(message)
	}	
}

//Returns the identified user list
func getUserList() []string {
	var listOfUsers []string
	for username, _ := range counter.users{
		listOfUsers = append(listOfUsers, username)
	}
	return listOfUsers
}

//gets the user received
func getUserProcessor(userName string)(*ServerProcessor, error){
	if userProcessor, ok := counter.users[userName]; ok {
		return userProcessor, nil
	}
	
	errorMessage := fmt.Sprintf("El usuario no existe '%s' ",
		userName)
	return nil, createError(errorMessage)
}


func getRoom(roomName string) (*room, error) {
	room, itExists := counter.rooms[roomName]
	if !itExists {
		errorString := fmt.Sprintf("El cuarto '%s' no existe",
			roomName)
		return nil, errors.New(errorString)
	}
	return room, nil
}

func verifyUserName(userName string) bool {
	if _, ok := counter.users[userName]; ok {
		return false
	}
	return true
}

//verifies if the status sent by client is valid.
func verifyStatus(status string) (bool) {
	switch status {
	case "AWAY": return true
	case "BUSY": return true
	case "ACTIVE": return true
	default: return false
	}
}

//verifies if the room name is available or not.
func verifyRoomName(roomName string) (bool) {
	if _, ok := counter.rooms[roomName]; ok {
		return false;
	}
	return true;
}

func verifyIdentifiedUsers(users []string) (bool, string) {
	for i := 0; i < len(users); i++ {
		if _, itExists := counter.users[users[i]]; !itExists {
			return false, users[i]
		}
	}
	return true, ""
}

func deleteUserName(oldName string) {
	delete(counter.users, oldName)
}

func deleteRoom(roomName string) {
	delete(counter.rooms, roomName)
}

func createNewRoom(host string, hostProcessor *ServerProcessor, roomName string) (error) {
	if value := strings.Compare(roomName, ""); value == 0 {
		return createError("Nombre de cuarto invalido")
	}
	if isRoomNameValid := verifyRoomName(roomName); isRoomNameValid {
		newRoom := getRoomInstance(roomName)
		newRoom.addUser(host)
		addRoom(roomName, newRoom)
		return nil
	}
	errorString := fmt.Sprintf("El cuarto '%s' ya existe",
		roomName)
	return createError(errorString)
}

func verifyRoomInvitation(host string, roomName string, usersToInvite []string) (error) {
	room, err := getRoom(roomName)
	
	if err != nil {
		errorString := fmt.Sprintf("El cuarto '%s' no existe", roomName)
		return errors.New(errorString)
	}
	if isMember := room.verifyRoomMember(host); !isMember {
		errorString := fmt.Sprintf("El usuario no esta en el cuarto '%s'", roomName)
		return createError(errorString)
	}
	
	return nil
}

func disconnectUserFromRoom(userName string, roomName string) (error) {
	room, error1 := getRoom(roomName)
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

func isEmptyRoom(roomName string) (bool) {
	room, _ := getRoom(roomName)
	return room.isEmpty()
}

func addInvitedUserToRoom(roomName string, userName string, userProcessor *ServerProcessor) {
	room, _ := getRoom(roomName)
	room.addInvitedUser(userName)
}

func addUserToRoom(userName string, roomName string, userProcessor *ServerProcessor) (error) {
	room, error := getRoom(roomName)
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

func removeInvitedUserInRoom(userName string, roomName string) {
	room, _ := getRoom(roomName)
	room.removeInvitedUser(userName)
}

func sendMessageToRoom(transmitter string, roomName string, message []byte) (error) {
	room, error:= getRoom(roomName)
	if  error != nil {
		return error
	}
	for userName, _ := range(room.users.elements) {
		if val := strings.Compare(userName, transmitter); val == 0 {
			continue
		}
		userProcessor, _ := getUserProcessor(userName)
		userProcessor.sendMessage(message)
	}
	return nil
}

func getRoomUserList(userName string, roomName string) ([]string, error) {
	room, err := getRoom(roomName)
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

