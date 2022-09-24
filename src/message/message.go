package message

import (
	"fmt"
	"encoding/json"
)

const (
	ERROR_MESSAGE_TYPE = "ERROR"
	WARNING_MESSAGE_TYPE = "WARNING"
	INFO_MESSAGE_TYPE = "INFO"
	IDENTIFY_MESSAGE_TYPE = "IDENTIFY"
	NEW_USER_TYPE = "NEW_USER"
)

type ErrorMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
}

type WarningMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Operation string `json:"operation"`
	Username string `json:"username"`
}

func GetWarningMessageJSON(mess WarningMessage) []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}


type InfoMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Operation string `json:"operation"`
}

func GetInfoMessageJSON(mess InfoMessage) []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type IdentifyMessage struct {
	Type string `json:"type"`
	Username string `json"username"`
}

type ChangeStatusMessage struct {
	Type string `json:"type"`
	Status string `json:"status"`
}

type SuccesChangeStatus struct {
	Type string `json:"type"`
	Status string `json:"status"`
	Operation string `json:"operation"`
}

type NewStatusMessage struct {
	Type string `json:"type"`
}

type NewUserMessage struct {
	Type string `json:"type"`
	Username string `json:"username"`
}

func GetNewUserMessageJSON(mess NewUserMessage) []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type StatusMessage struct {	
	Type string `json:"type"`
}

type GetUsersMessage struct {
	Type string `json:"type"`
}

