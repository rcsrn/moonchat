package message

import (
	"fmt"
	"encoding/json"
)

const (
	PUBLIC_MESSAGE_FROM_TYPE = "PUBLIC_MESSAGE_FROM" 
	ERROR_TYPE = "ERROR"
	WARNING_TYPE = "WARNING"
	INFO_TYPE = "INFO"
	NEW_USER_TYPE = "NEW_USER"
	DISCONNECTED_TYPE = "DISCONNECTED"
	NEW_STATUS_TYPE = "NEW_STATUS"
	USER_LIST_TYPE = "USER_LIST"
	PRIVATE_TYPE = "MESSAGE_FROM"
	INVITATION_TYPE = "INVITATION"
	JOINED_ROOM_TYPE = "JOINED_ROOM_TYPE"
	ROOM_USER_LIST_TYPE = "ROOM_USER_LIST"
	ROOM_MESSAGE_FROM_TYPE = "ROOM_MESSAGE_FROM"
	LEFT_ROOM_TYPE = "LEFT_ROOM_TYPE"
)

type Message interface {
	GetJSON() []byte
}

type IdentifyErrorMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
}

func (mess IdentifyErrorMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type StatusErrorMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Operation string `json:operation"`
}

func (mess StatusErrorMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type RoomErrorMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Operation string `json:operation"`
	Roomname string `json:roomname`
}

func (mess RoomErrorMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type RoomWarningMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Operation string `json:"operation"`
	Roomname string `json:"roomname"`
}

func (mess RoomWarningMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type UsernameWarningMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Operation string `json:"operation"`
	Username string `json:"username"`
}

func (mess UsernameWarningMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type StatusWarningMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Operation string `json:"operation"`
	Status string `json:"status"`
}

func (mess StatusWarningMessage) GetJSON() []byte {
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

type RoomSuccesMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Operation string `json:"operation"`
	Roomname string `json:"roomname"`
}

func (mess RoomSuccesMessage) GetJSON() []byte {
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
	Status string `json:"status"`
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
	Type string `json:"type"`
	Usernames []string `json:"usernames"`	
}

func (mess UserList) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type NewMessage struct {
	Type string `json:"type"`
	Username string `json:"username"`
	Message string `json:"message"`
}

func (mess NewMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type LeftRoomMessage struct {
	Type string `json:"type"`
	Roomname string `json:"roomname"`
	Username string `json:"username"`	
}

func (mess LeftRoomMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type RoomMessage struct {
	Type string `json:"type"`
	Roomname string `json:"roomname"`
	Username string `json:"username"`
	Message string `json:"message"`
}

func (mess RoomMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}
