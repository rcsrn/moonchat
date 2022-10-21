package view

import (
	"bufio"
	"log"
	"strings"
	"os"
)

type ConsoleListener struct {
	buffer []string
	isReadBuffer bool
}

func GetConsoleListenerInstance() *ConsoleListener {
	return &ConsoleListener{make([]string, 512), false}
}

func (listener *ConsoleListener) ListenFromConsole() []string {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		listener.buffer = append(listener.buffer, line)
		return listener.buffer
	}
	return nil
}
