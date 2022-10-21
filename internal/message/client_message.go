package message

import (
	"fmt"
	"encoding/json"
)

const (
	PUBLIC_TYPE = "PUBLIC_MESSAGE"
	IDENTIFY_TYPE = "IDENTIFY"
	DISCONNECT_TYPE = "DISCONNECT"
	STATUS_TYPE = "STATUS"
	USERS_TYPE = "USERS"
	MESSAGE_TYPE = "MESSAGE"
	NEW_ROOM_TYPE = "NEW_ROOM"
	INVITE_TYPE =  "INVITE"
	JOIN_ROOM_TYPE = "JOIN_ROOM"
	ROOM_USERS_TYPE = "ROOM_USERS"
	ROOM_MESSAGE_TYPE = "ROOM_MESSAGE"
	LEAVE_ROOM_TYPE = "LEAVE_ROOM"
)

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


