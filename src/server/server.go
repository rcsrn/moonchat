package main

import(
	"net"
	"fmt"
	"github.com/rcsrn/moonchat/src/message"
	//"encoding/json"
)

type Server struct {	
}

var users = make(map[string]ServerProcessor)


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

func checkIdentify(username string) []byte {
	if _, ok := users[username]; ok {
		m := message.WarningMessage{message.WARNING_MESSAGE_TYPE, "username already used" , message.IDENTIFY_MESSAGE_TYPE, username}
		return message.GetWarningMessageJSON(m)
	}
	m := message.InfoMessage{message.INFO_MESSAGE_TYPE, "Succes: username has been saved", message.IDENTIFY_MESSAGE_TYPE}
	return message.GetInfoMessageJSON(m)
}




