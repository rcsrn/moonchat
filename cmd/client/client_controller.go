package main

import (
	"github.com/rcsrn/moonchat/internal/client"
	"github.com/rcsrn/moonchat/internal/view"
	"os"
	"fmt"
)

var printer *view.Printer
var listener *view.ConsoleListener

func main () {
	printer = view.GetPrinterInstance()

	fmt.Println(len(os.Args))
	fmt.Println(os.Args)
	if len(os.Args) != 3  {
		printer.Use()
		os.Exit(1)
	}

	host := getHostDirection(os.Args)
	
	fmt.Println(host)
	client := client.GetClientInstance()
	error := client.Connect(host)
	
	if error != nil {
		printer.WarnConnectError()
		os.Exit(1)
	}
	printer.PrintInstructions()
	listener = view.GetConsoleListenerInstance()
	

	for {		
	message := listener.ListenFromConsole()
	client.ProcessMessage(message)
	}
}

func getHostDirection(arguments []string) string {
	return arguments[1] + ":" + arguments[2]
}