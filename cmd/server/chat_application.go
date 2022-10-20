package main

type application interface {
	Start() 
}

type chat struct {
	server *server
}

func GetChatInstance() *chat {
	chat := chat{&server{}}
	return &chat
}

func (chat *chat) Start() {
	chat.server.initServer()
	chat.server.waitForConnections()
}


