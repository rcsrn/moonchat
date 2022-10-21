package view

import (
	//"bufio"
)

type ConsoleListener struct {
	buffer []string
}

func GetConsoleListenerInstance() *ConsoleListener {
	return nil
}

func (listener *ConsoleListener) ListenFromConsole() {
	// reader := bufio.NewReader(os.Stdin)

	// var readLines []string
	// for {
	// 	line, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if len(strings.TrimSpace(line)) == 0 {
	// 		break
	// 	}
	// 	readLines = append(readLines, line)
	// }
}

func (listener *ConsoleListener) GetBuffer() []string {
	return listener.buffer
}
