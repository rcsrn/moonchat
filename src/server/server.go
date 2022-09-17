package main

import(
	"net"
	"fmt"
	"github.com/rcsrn/moonchat/src/message"
)

type Server struct {
}

var users = make(map[string]*ServerProcessor)

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
		serverProcessor := ServerProcessor{connection}
		go serverProcessor.readMessages()
	}
}

func checkIdentify(username string, processor *ServerProcessor) []byte {
	if _, ok := users[username]; ok {
		m := message.WarningMessage{message.WARNING_MESSAGE_TYPE, "username already used" , message.IDENTIFY_MESSAGE_TYPE, username}
		return message.GetWarningMessageJSON(m)
	}
	addUser(username, processor)
	m := message.InfoMessage{message.INFO_MESSAGE_TYPE, "Succes: username has been saved", message.IDENTIFY_MESSAGE_TYPE}
	return message.GetInfoMessageJSON(m)
}

func addUser(username string, processor *ServerProcessor) {
	users[username] = processor
	m := message.NewUserMessage(message.NEW_USER_TYPE, username)
	
	//toAllUsers(mesage.GetNewUserMessageJSON(m))
}
// Debe iterar el diccionario de usuarios e ir diciendo a los procesadores que envien ese mensaje a traves de cada conexion.
//func toAllUsers(message []byte) {
	
//}

