package server

type application interface {
	Start() 
}

type chat struct {
	server *Server
}

func GetChatInstance() *chat {
	chat := chat{}
	chat.server = GetServerInstance()
	return &chat
}

func (chat *chat) Start() {
	chat.server.waitForConnections()
}


