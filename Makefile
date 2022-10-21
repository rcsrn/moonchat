buildServer: ./cmd/chat
	go build ./cmd/chat/*.go

startServer:
	./chat

buildClient : ./cmd/client
	go build ./cmd/client/*.go

startClient:
	./client_controller
