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
	NEW_USER_MESSAGE_TYPE = "NEW_USER"
	DISCONNECTED_MESSAGE_TYPE = "DISCONNECTED"
	DISCONNECT_MESSAGE_TYPE = "DISCONNECT"
	STATUS_MESSAGE_TYPE = "STATUS"
	NEW_STATUS_MESSAGE_TYPE = "NEW_STATUS"
	USERS_MESSAGE_TYPE = "USERS"
	USER_LIST_MESSAGE_TYPE = "USER_LIST"
	MESSAGE_TYPE = "MESSAGE"
)

type Message interface {
	GetJSON() []byte
}


type ErrorMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
}

type WarningMessageUsername struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Operation string `json:"operation"`
	Username string `json:"username"`
}

func (mess WarningMessageUsername) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type WarningMessageStatus struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Operation string `json:"operation"`
	Status string `json:"status"`
}

func (mess WarningMessageStatus) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type SuccesMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Operation string `json:"operation"`
}

func (mess SuccesMessage) GetJSON() []byte {
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

func (mess IdentifyMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type NewUserMessage struct {
	Type string `json:"type"`
	Username string `json:"username"`
}

func (mess NewUserMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type NewStatusMessage struct {
	Type string `json:"type"`
	Username string `json:"username"`
	Status string `json:status`
}

func (mess NewStatusMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type DisconnectedMessage struct {
	Type string `json:"type"`
	Username string `json:"username"`
}

func (mess DisconnectedMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type UserList struct {
	Type string `json:type`
	Usernames []string `json:username`	
}

func (mess UserList) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type NewMessage struct {
	Type string `json:type`
	Username string `json:username`
	Message string `json:message`
}

func (mess NewMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}


