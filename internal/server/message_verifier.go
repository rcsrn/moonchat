package server

import (
	"github.com/rcsrn/moonchat/internal/message"
	//"fmt"
)

type messageVerifier struct {
	message map[string]string
}

func (verifier *messageVerifier) setMessage(message map[string]string) {
	verifier.message = message
}

func (verifier *messageVerifier) validateMessage() (bool) {
	switch verifier.message["type"] {
	case message.IDENTIFY_TYPE:
		return VerifyIdentifyMessage(verifier.message)
	case message.STATUS_TYPE:
		return VerifyStatusMessage(verifier.message)
	case message.USERS_TYPE:
		return VerifyUserMessage(verifier.message)
	case message.DISCONNECT_TYPE:
		return VerifyDisconnectMessage(verifier.message)
	case message.PUBLIC_TYPE:
		return VerifyPublicMessage(verifier.message)
	case message.MESSAGE_TYPE:
		return VerifyPrivateMessage(verifier.message)
	case message.NEW_ROOM_TYPE:
		return VerifyNewRoomMessage(verifier.message)
	case message.INVITE_TYPE:
		return VerifyInviteMessage(verifier.message)
	case message.JOIN_ROOM_TYPE:
		return VerifyJoinMessage(verifier.message)
	case message.ROOM_USERS_TYPE:
		return VerifyRoomUserMessage(verifier.message)
	case message.ROOM_MESSAGE_TYPE:
		return VerifyRoomMessage(verifier.message)
	case message.LEAVE_ROOM_TYPE:
		return VerifyLeaveMessage(verifier.message)
	default: return false	
	}
} 

func VerifyIdentifyMessage(identifyMessage map[string]string) (bool) {
	if len(identifyMessage) != 2 {
		return false
	}
	if _, itExists := identifyMessage["type"]; !itExists {
		return false
	}
	if _, itExists := identifyMessage["username"]; !itExists {
		return false
	}
	return true
}

func VerifyStatusMessage(statusMessage map[string]string) (bool) {
	if len(statusMessage) != 2 {
		return false
	}
	if _, itExists := statusMessage["type"]; !itExists {
		return false
	}
	if _, itExists := statusMessage["status"]; !itExists {
		return false
	}
	return true
}

func VerifyUserMessage(userMessage map[string]string) (bool) {
	if len(userMessage) != 1 {
		return false
	}
	if _, itExists := userMessage["type"]; !itExists {
		return false
	}
	return true

}

func VerifyDisconnectMessage(disconnectMessage map[string]string) (bool) {
	if len(disconnectMessage) != 1 {
		return false
	}
	if _, itExists := disconnectMessage["type"]; !itExists {
		return false
	}
	return true

}

func  VerifyPublicMessage(publicMessage map[string]string) (bool) {
	if len(publicMessage) != 2 {
		return false
	}
	if _, itExists := publicMessage["type"]; !itExists {
		return false
	}
	if _, itExists := publicMessage["message"]; !itExists {
		return false
	}
	return true

}

func VerifyPrivateMessage(privateMessage map[string]string) (bool) {
	if len(privateMessage) != 3 {
		return false
	}
	if _, itExists := privateMessage["type"]; !itExists {
		return false
	}
	if _, itExists := privateMessage["username"]; !itExists {
		return false
	}
	if _, itExists := privateMessage["message"]; !itExists {
		return false
	}
	return true
}

func VerifyNewRoomMessage(newRoomMessage map[string]string) (bool) {
	if len(newRoomMessage) != 2 {
		return false
	}
	if _, itExists := newRoomMessage["type"]; !itExists {
		return false
	}
	if _, itExists := newRoomMessage["roomname"]; !itExists {
		return false
	}
	return true

}

func VerifyInviteMessage(inviteMessage map[string]string) (bool) {
	if len(inviteMessage) != 3 {
		return false
	}
	if _, itExists := inviteMessage["type"]; !itExists {
		return false
	}
	if _, itExists := inviteMessage["roomname"]; !itExists {
		return false
	}
	if _, itExists := inviteMessage["usernames"]; !itExists {
		return false
	}
	return true

}

func VerifyJoinMessage(joinMessage map[string]string) (bool) {
	if len(joinMessage) != 2 {
		return false
	}
	if _, itExists := joinMessage["type"]; !itExists {
		return false
	}
	if _, itExists := joinMessage["roomname"]; !itExists {
		return false
	}
	return true

}

func VerifyRoomUserMessage(roomUserMessage map[string]string) (bool) {
	if len(roomUserMessage) != 2 {
		return false
	}
	if _, itExists := roomUserMessage["type"]; !itExists {
		return false
	}
	if _, itExists := roomUserMessage["roomname"]; !itExists {
		return false
	}
	return true

}

func VerifyRoomMessage(roomMessage map[string]string) (bool) {
	if len(roomMessage) != 3 {
		return false
	}
	if _, itExists := roomMessage["type"]; !itExists {
		return false
	}
	if _, itExists := roomMessage["roomname"]; !itExists {
		return false
	}
	if _, itExists := roomMessage["message"]; !itExists {
		return false
	}
	return true

}

func VerifyLeaveMessage(leaveMessage map[string]string) (bool) {
	if len(leaveMessage) != 2 {
		return false
	}
	if _, itExists := leaveMessage["type"]; !itExists {
		return false
	}
	if _, itExists := leaveMessage["roomname"]; !itExists {
		return false
	}
	return true

}
