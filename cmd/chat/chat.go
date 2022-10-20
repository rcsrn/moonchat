package main

import (
	"github.com/rcsrn/moonchat/internal/server"
)

func main() {
	chat := server.GetChatInstance()
	chat.Start();
}
