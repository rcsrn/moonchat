package main

type chatApplication struct {
	server server
}


func GetChatApplicationInstance() *chatApplication {
	chat := chatApplication{}
	return &chat
}

func (chat *chatApplication) Start() {
	chat.server = server{}
	chat.server.initServer()
	chat.server.waitForConnections()
}


