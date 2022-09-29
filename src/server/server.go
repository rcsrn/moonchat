package main

import(
	"net"
	"fmt"
	"github.com/rcsrn/moonchat/src/message"
	"sync"
	"errors"
)

type Server struct {
}

//Map protected for concurrency
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
	fmt.Println("Server is already")
	connectionListener, err := net.Listen(SERVER_TYPE, SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for connections...")
	if err != nil {
		fmt.Println("Algo ha salido mal. :(")
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

func checkIdentify(username string, processor *ServerProcessor) []byte {
	counter.blocker.RLock()
	if _, ok := counter.users[username]; ok {
		m := message.WarningMessageUsername{message.WARNING_MESSAGE_TYPE, "username already used" , message.IDENTIFY_MESSAGE_TYPE, username}
		return m.GetJSON()
	}
	counter.blocker.RUnlock()
	addUser(username, processor)
	processor.setUserName(username)
	m := message.SuccesMessage{message.INFO_MESSAGE_TYPE, "Succes: username has been saved", message.IDENTIFY_MESSAGE_TYPE}
	return m.GetJSON()
}

func addUser(username string, processor *ServerProcessor) {
	//counter.blocker.Lock()
	counter.users[username] = processor
	//counter.blocker.Unlock()
	m := message.NewUserMessage{message.NEW_USER_MESSAGE_TYPE, username}
	toAllUsers(m.GetJSON())
}

func addRoom(roomname string, room *room) {
	counter.blocker.Lock()
	defer counter.blocker.Unlock()
	counter.rooms["roomname"] = room
}

func toAllUsers(message []byte) {
	counter.blocker.RLock()
	for _, element := range counter.users {
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

func sendPrivateMessage(receiver string, messageToSend string, transmitter string) (error){
	userProcess, err := getUserProcessor(receiver)
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
	if userProcessor, ok := counter.users["userName"]; ok {
		return userProcessor, nil
	}
	return nil, errors.New("User not found")
}

//verifies the statatus sent by client.x
func verifyStatus(status string) (bool) {
	switch status {
	case "AWAY": return true
	case "BUSY": return true
	case "ACTIVE": return true
	default: return false
	}
}

func verifyRoomName(roomName string) (bool) {
	counter.blocker.RLock()
	defer counter.blocker.RUnlock()
	if _, ok := counter.rooms["roomName"]; ok {
		return false;
	}
	return true;
}

func createNewRoom(host string, hostProcessor *ServerProcessor, roomname string) ([]byte) {
	if ok := verifyRoomName(host); ok {
		var users map[string]*ServerProcessor
		users["host"] = hostProcessor
		var blocker sync.RWMutex = sync.RWMutex{}
		var roomUsers mapCounter = mapCounter{blocker,users}
		var newRoom room = room{roomUsers, roomname}
		addRoom(roomname, &newRoom)
		succes := message.SuccesMessage{message.INFO_MESSAGE_TYPE, "Succes: The room has been created succesfully.", message.NEW_ROOM_MESSAGE_TYPE}
		return succes.GetJSON()
	}
	str := fmt.Sprintf("The '%s' already exists.", roomname)
	fail := message.RoomWarningMessage{message.WARNING_MESSAGE_TYPE, str, message.NEW_ROOM_MESSAGE_TYPE, roomname}
	return fail.GetJSON()
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

