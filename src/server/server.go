package main

import(
	"net"
	"fmt"
	"github.com/rcsrn/moonchat/src/message"
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

func (server *server) WaitForConnections() {
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
		serverProcessor := ServerProcessor{connection, "", "ACTIVE", false}
		go serverProcessor.readMessages()
	}
}

func (server *server) initRooms() {
	counter.rooms = make(map[string]*room)
}

func addUser(userName string, processor *ServerProcessor) {
	counter.users[userName] = processor
	m := message.NewUserMessage{message.NEW_USER_TYPE, userName}
	toAllUsers(processor, m.GetJSON())
}

func addRoom(roomname string, room *room) {
	counter.rooms[roomname] = room
}

//sends a message just to users that have been added.
func toAllUsers(processor *ServerProcessor, message []byte) {
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
	
	errorMessage := fmt.Sprintf("The user '%s' does not exist.",
		userName)
	return nil, createError(errorMessage)
}


func getRoom(roomName string) (*room, error) {
	room, itExists := counter.rooms[roomName]
	if !itExists {
		errorString := fmt.Sprintf("The room '%s' does not exist.",
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

func removeOldName(oldName string) {
	delete(counter.users, oldName)
}

func createNewRoom(host string, hostProcessor *ServerProcessor, roomName string) (error) {
	if value := strings.Compare(roomName, ""); value == 0 {
		return createError("Invalid roomName!")
	}
	if isRoomNameValid := verifyRoomName(roomName); isRoomNameValid {
		newRoom := getRoomInstance(roomName)
		newRoom.init()
		newRoom.addUser(host, hostProcessor)
		addRoom(roomName, newRoom)
		return nil
	}
	errorString := fmt.Sprintf("Roomname '%s' is already used.",
		roomName)
	return createError(errorString)
}

func verifyRoomInvitation(host string, roomName string, usersToInvite []string) (error) {
	room, err := getRoom(roomName)
	
	if err != nil {
		errorString := fmt.Sprintf("The room '%s' does not exist.", roomName)
		return errors.New(errorString)
	}
	if isMember := room.verifyRoomMember(host); !isMember {
		errorString := fmt.Sprintf("The user is not a member of the room '%s'.", roomName)
		return errors.New(errorString)
	}
	
	return nil
}


func addInvitedUser(roomName string, userName string, userProcessor *ServerProcessor) {
	room, _ := getRoom(roomName)
	room.addInvitedUser(userName, userProcessor)
}

func addUserToRoom(userName string, roomName string, userProcessor *ServerProcessor) (error) {
	room, error := getRoom(roomName)
	if error != nil {
		return error
	}
	if userHasBeenInvited := room.verifyInvitedUser(userName); !userHasBeenInvited {
		errorMessage := fmt.Sprintf("The user has not been invited to '%s'.",
			roomName)
		return createError(errorMessage)
	}
	room.addUser(userName, userProcessor)
	return nil
}

func toAllRoomUsers(transmitter string, roomName string, message []byte) {
	room, _:= getRoom(roomName)
	room.sendToAllUsers(transmitter, message)
}

func getRoomUserList(userName string, roomName string) ([]string, error) {
	room, err := getRoom(roomName)
	if err != nil {
		errorString := fmt.Sprintf("The room '%s' does not exist.",
			roomName)
		return nil, errors.New(errorString)
	}
	if isMember := room.verifyRoomMember(userName); !isMember {
		errorString := fmt.Sprintf("The user '%s' is not a member of the room '%s'",
			userName,
			roomName)
		return nil , errors.New(errorString)
	}
	return room.getMemberList(), nil
}

func createError(errorMessage string) (error) {
	return errors.New(errorMessage)
}

