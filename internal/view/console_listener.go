package view

import (
	"bufio"
	"log"
	"strings"
	"os"
)

type ConsoleListener struct {
}

func GetConsoleListenerInstance() *ConsoleListener {
	return &ConsoleListener{}
}

func (listener *ConsoleListener) ListenFromConsole() []string {
	reader := bufio.NewReader(os.Stdin)

	var buffer []string
	
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		buffer = append(buffer, line)
		return buffer
	}
	return nil
}
