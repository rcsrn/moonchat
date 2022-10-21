package message

import (
	"fmt"
	"encoding/json"
)

const (
	PUBLIC_MESSAGE_FROM_TYPE = "PUBLIC_MESSAGE_FROM" 
	PUBLIC_TYPE = "PUBLIC_MESSAGE"
	ERROR_TYPE = "ERROR"
	WARNING_TYPE = "WARNING"
	INFO_TYPE = "INFO"
	IDENTIFY_TYPE = "IDENTIFY"
	NEW_USER_TYPE = "NEW_USER"
	DISCONNECTED_TYPE = "DISCONNECTED"
	DISCONNECT_TYPE = "DISCONNECT"
	STATUS_TYPE = "STATUS"
	NEW_STATUS_TYPE = "NEW_STATUS"
	USERS_TYPE = "USERS"
	USER_LIST_TYPE = "USER_LIST"
	MESSAGE_TYPE = "MESSAGE"
	PRIVATE_TYPE = "MESSAGE_FROM"
	NEW_ROOM_TYPE = "NEW_ROOM"
	INVITE_TYPE =  "INVITE"
	INVITATION_TYPE = "INVITATION"
	JOIN_ROOM_TYPE = "JOIN_ROOM"
	JOINED_ROOM_TYPE = "JOINED_ROOM_TYPE"
	ROOM_USERS_TYPE = "ROOM_USERS"
	ROOM_USER_LIST_TYPE = "ROOM_USER_LIST"
	ROOM_MESSAGE_TYPE = "ROOM_MESSAGE"
	ROOM_MESSAGE_FROM_TYPE = "ROOM_MESSAGE_FROM"
	LEAVE_ROOM_TYPE = "LEAVE_ROOM"
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


//Client Messages:

type IdentifyMessage struct {
	Type string `json:"type"`
	Username string `json:"username"`
}

func (mess IdentifyMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type InviteToRoomMessage struct {
	Type string `json:"type"`
	Roomname string `json:"roomname"`
	Usernames []string `json:"usernames"`
}

func (mess InviteToRoomMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type RoomInvitationMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Username string `json:"username"`
	Roomname string `json:"roomname"`
}

func (mess RoomInvitationMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type JoinedRoomMessage struct {
	Type string `json:"type"`
	Roomname string `json:"roomname"`
	Username string `json:"username"`
}

func (mess JoinedRoomMessage) GetJSON() []byte {
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

type UserMessage struct {
	Type string `json:"type"`
}

func (mess UserMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type StatusMessage struct {
	Type string `json:"type"`
	Status string `json:"status"`
}

func (mess StatusMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type DisconnectMessage struct {
	Type string `json:"type"`
}

func (mess DisconnectMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type PublicMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
}

func (mess PublicMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type PrivateMessage struct {
	Type string `json:"type"`
	Username string `json:"username"`
	Message string `json:"message"`
}

func (mess PrivateMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type NewRoomMessage struct {
	Type string `json:"type"`
	Roomname string `json:"roomname"`
	Message string `json:"message"`
}

func (mess NewRoomMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type JoinMessage struct {
	Type string `json:"type"`
	Roomname string `json:"roomname"`
}

func (mess JoinMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type RoomUserMessage struct {
	Type string `json:"type"`
	Roomname string `json:"roomname"`
}

func (mess RoomUserMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

type LeaveMessage struct {
	Type string `json:"type"`
	Roomname string `json:"roomname"`
}

func (mess LeaveMessage) GetJSON() []byte {
	message, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("This should not happen")
	}
	return message
}

