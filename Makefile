buildServer: ./src/server
	go build ./src/server/*.go

startServer:
	./chat_application
