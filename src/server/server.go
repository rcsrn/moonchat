package main

import(
	"net"
	"fmt"
	"github.com/rcsrn/moonchat/src/message"
	"encoding/json"
)

type Server struct {
}

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "1234"
	SERVER_TYPE = "tcp"
)

func (s *Server) WaitForConnections() {
	connectionListener, err := net.Listen(SERVER_TYPE, SERVER_HOST + ":" + SERVER_PORT)
	if err != nil {
		fmt.Println("Algo ha salido mal. :(")
	}
	
	for {
		connection, err := connectionListener.Accept()
		if err != nil {
			fmt.Println("FAIL: Connection denied.")
			continue
		}
		serverProcessor := ServerProcessor{connection}
		serverProcessor.sendToClient(connectionAccepted())
		go serverProcessor.processClient()
	}
}

//Sends a message that connection was accepted.
func connectionAccepted() []byte {
	message := message.GetMessage(1)	
	return message
}

//Decoding of message sent by Client.
func readMessage(b []byte) {
	var message string
	json.Unmarshal(b, &message)
	fmt.Println(message)
}
