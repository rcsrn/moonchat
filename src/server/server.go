package main

import(
	"net"
	"fmt"
	"github.com/rcsrn/moonchat/src/message"
	"sync"
	"errors"
	"strings"
)

type Server struct {
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

func (server *Server) WaitForConnections() {
	initRooms()
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
		serverProcessor := ServerProcessor{connection, "", "ACTIVE"}
		go serverProcessor.readMessages()
	}
}

func initRooms() {
	counter.rooms = make(map[string]*room)
}

func addUser(username string, processor *ServerProcessor) {
	counter.blocker.Lock()
	counter.users[username] = processor
	counter.blocker.Unlock()
	m := message.NewUserMessage{message.NEW_USER_MESSAGE_TYPE, username}
	toAllUsers(processor, m.GetJSON())
}

func addRoom(roomname string, room *room) {
	//counter.blocker.Lock()
	//defer counter.blocker.Unlock()
	counter.rooms[roomname] = room
}

//sends a message just to users that have been added.
func toAllUsers(processor *ServerProcessor, message []byte) {
	counter.blocker.RLock()
	for _, element := range counter.users {
		if processor == element {
			continue
		}
		element.sendMessage(message)
	}	
	counter.blocker.RUnlock()
}

//Returns the identified user list
func getUserList() []byte {
	counter.blocker.RLock()
	var listOfUsers []string
	for username, _ := range counter.users{
		listOfUsers = append(listOfUsers, username)
	}
	counter.blocker.RUnlock()
	mess := message.UserList{message.USER_LIST_MESSAGE_TYPE, listOfUsers}
	return mess.GetJSON()
}

func sendPrivateMessage(receptor string, messageToSend string, transmitter string) (error){
	userProcess, err := getUserProcessor(receptor)
	if err != nil {
		return err
	}
	privateMessage := message.NewMessage{message.PRIVATE_MESSAGE_TYPE, transmitter, messageToSend}
	userProcess.sendMessage(privateMessage.GetJSON())
	return nil
}

//gets the user received
func getUserProcessor(userName string)(*ServerProcessor, error){
	counter.blocker.RLock()
	defer counter.blocker.RUnlock()
	if userProcessor, ok := counter.users[userName]; ok {
		return userProcessor, nil
	}
	return nil, errors.New("User not found")
}

func verifyUserName(userName string) bool {
	counter.blocker.RLock()
	if _, ok := counter.users[userName]; ok {
		return false
	}
	counter.blocker.RUnlock()
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
	counter.blocker.RLock()
	defer counter.blocker.RUnlock()
	if _, ok := counter.rooms[roomName]; ok {
		return false;
	}
	return true;
}


func createNewRoom(host string, hostProcessor *ServerProcessor, roomname string) ([]byte, error) {
	if value := strings.Compare(roomname, ""); value == 0 {
		return nil, errors.New("Invalid room name")
	}
	if ok := verifyRoomName(roomname); ok {
		var users map[string]*ServerProcessor = make(map[string]*ServerProcessor)
		users["host"] = hostProcessor
		var blocker sync.RWMutex = sync.RWMutex{}
		var roomUsers mapCounter = mapCounter{blocker,users}
		var newRoom room = room{roomUsers, roomname}
		addRoom(roomname, &newRoom)
		str := fmt.Sprintf("Succes: The room '%s'has been created succesfully.", roomname)
		succes := message.SuccesMessage{message.INFO_MESSAGE_TYPE, str, message.NEW_ROOM_MESSAGE_TYPE}
		return succes.GetJSON(), nil
	}
	str := fmt.Sprintf("The '%s' already exists.", roomname)
	fail := message.RoomWarningMessage{message.WARNING_MESSAGE_TYPE, str, message.NEW_ROOM_MESSAGE_TYPE, roomname}
	return fail.GetJSON(), errors.New("Room name already used")
}

func inviteUsersToRoom(roomName string, users string) []byte {
	return nil
}

func joinRoom(username string, roomName string) []byte {
	return nil
}

func getRoomUserList(username string, roomName string) []byte {
	return nil
}

