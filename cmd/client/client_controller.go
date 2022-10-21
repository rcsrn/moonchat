package main

import (
	"github.com/rcsrn/moonchat/internal/client"
	"github.com/rcsrn/moonchat/internal/view"
	"os"
)

var printer *view.Printer
var listener *view.ConsoleListener

func main () {
	client := client.GetClientInstance()
	
	error := client.Connect()
	
	if error != nil {
		printer.WarnConnectError()
		os.Exit(1)
	}
	
	printer = view.GetPrinterInstance()
	listener = view.GetConsoleListenerInstance()

	go getBufferOfMessages(client)
	go listener.ListenFromConsole()

	printer.PrintInstructions()
	
	for !client.IsIdentified() {
		printer.RequestUserName()
	}
}

func getBufferOfMessages(client *client.Client) {
	for {
		buffer := listener.GetBuffer()
		if buffer != nil {
			client.ProcessMessage(buffer)
		}
	}
}

