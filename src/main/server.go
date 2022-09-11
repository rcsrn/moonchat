package main

import (
	"fmt"
	"net"
)

type server struct {
	
}

//Start running the server
func (server *server) createServer() (Listener ,error) {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		return nil,err
	}
	return listener, nil
}


func (server *server) waitForConnections(listener Listener) {
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Esto debe enviar un mensaje de que la peticion de conexion ha sido rechazada.")
			continue
		} else {
			fmt.Println("Se debe enviar un mensaje de que la peticion de conexion ha sido aceptada.")
		}
	}
}
