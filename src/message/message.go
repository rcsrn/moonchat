package message

import (
	"fmt"
	"encoding/json"
)

const(	
	id string = "IDENTIFY"
	info string = "INFO"
	warning string = "WARNING"
	error string = "ERROR"
)

type ErrorMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
}

func GetMessage(num int) []byte {
	switch (num) {
	case 0: message := ErrorMessage{error, "Mensaje de error"}
		return getJson(message)
	}
	return nil
}

func getJson(message interface{}) []byte {
	json, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Esto no deberia ocurrir")
	}
	return json
}


