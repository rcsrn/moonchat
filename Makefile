buildServer: ./cmd/chat
	go build ./cmd/chat/*.go

startServer:
	./chat
