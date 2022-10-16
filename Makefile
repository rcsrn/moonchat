buildServer: ./cmd/server
	go build ./cmd/server/*.go

startServer:
	./chat_application
