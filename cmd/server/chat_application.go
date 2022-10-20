package main

type chatAplication struct {
	server server
}


func (chat *chatAplication)start() {
	chat.server = server{}
	chat.server.initServer()
	chat.server.waitForConnections()
}
