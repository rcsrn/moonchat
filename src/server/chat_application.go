package main

type chatAplication struct {
	server server
}


func (chat *chatAplication)start() {
	chat.server = server{}
	chat.server.initRooms()
	chat.server.WaitForConnections()
}
