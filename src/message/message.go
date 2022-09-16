package message

// import (
// 	"fmt"
// 	"encoding/json"
// )

const(
	
	id string = "IDENTIFY"	
)


func GetMessage(typeOfMessage int) []byte {
	return nil
}

type ErrorMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
}

// func getJson(message Message) []byte {
// 	json, err := json.Marshal(message)
// 	if err != nil {
// 		fmt.Println("Esto no deberia ocurrir")
// 	}
// 	return json
// }


