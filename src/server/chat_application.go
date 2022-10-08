package main

type chatAplication struct {
	
}

func (chat *chatAplication)start() {
	server := Server{}
	initRooms()
	server.WaitForConnections()
}
