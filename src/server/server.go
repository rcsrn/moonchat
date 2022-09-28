package main

import(
	"net"
	"fmt"
	"github.com/rcsrn/moonchat/src/message"
	"sync"
)

type Server struct {
}

//Map protected for concurrency
var counter = struct{
    sync.RWMutex
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
		//It is neccesary to add an default status.
		serverProcessor := ServerProcessor{connection, "", ""}
		go serverProcessor.readMessages()
	}
}

func checkIdentify(username string, processor *ServerProcessor) []byte {
	counter.RLock()
	if _, ok := counter.users[username]; ok {
		m := message.WarningMessageUsername{message.WARNING_MESSAGE_TYPE, "username already used" , message.IDENTIFY_MESSAGE_TYPE, username}
		return m.GetJSON()
	}
	counter.RUnlock()
	addUser(username, processor)
	processor.setUserName(username)
	m := message.SuccesMessage{message.INFO_MESSAGE_TYPE, "Succes: username has been saved", message.IDENTIFY_MESSAGE_TYPE}
	return m.GetJSON()
}

func addUser(username string, processor *ServerProcessor) {
	counter.Lock()
	counter.users[username] = processor
	counter.Unlock()
	m := message.NewUserMessage{message.NEW_USER_MESSAGE_TYPE, username}
	toAllUsers(m.GetJSON())
}

func toAllUsers(message []byte) {
	counter.RLock()
	for _, element := range counter.users {
		element.sendMessage(message)
	}	
	counter.RUnlock()
	fmt.Println("SI LLEGA AQUI SI TERMINO")	
}

//Returns the identified user list
func getUserList() []byte {
	var listOfUsers []string
	for username, _ := range counter.users{
		listOfUsers = append(listOfUsers, username)
	}
	mess := message.UserList{message.USER_LIST_MESSAGE_TYPE, listOfUsers}
	return mess.GetJSON()
}

//verifies the statatus sent by client.x
func verifyStatus (status string) bool {
	switch status {
	case "AWAY": return true
	case "BUSY": return true
	default: return false
	}
}


