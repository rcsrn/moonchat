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
		return verifyIdentifyMessage(verifier.message)
	case message.STATUS_TYPE:
		return verifyStatusMessage(verifier.message)
	case message.USERS_TYPE:
		return verifyUserMessage(verifier.message)
	case message.DISCONNECT_TYPE:
		return verifyDisconnectMessage(verifier.message)
	case message.PUBLIC_TYPE:
		return verifyPublicMessage(verifier.message)
	case message.MESSAGE_TYPE:
		return verifyPrivateMessage(verifier.message)
	case message.NEW_ROOM_TYPE:
		return verifyNewRoomMessage(verifier.message)
	case message.INVITE_TYPE:
		return verifyInviteMessage(verifier.message)
	case message.JOIN_ROOM_TYPE:
		return verifyJoinMessage(verifier.message)
	case message.ROOM_USERS_TYPE:
		return verifyRoomUserMessage(verifier.message)
	case message.ROOM_MESSAGE_TYPE:
		return verifyRoomMessage(verifier.message)
	case message.LEAVE_ROOM_TYPE:
		return verifyLeaveMessage(verifier.message)
	default: return false	
	}
} 

func verifyIdentifyMessage(identifyMessage map[string]string) (bool) {
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

func verifyStatusMessage(statusMessage map[string]string) (bool) {
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

func verifyUserMessage(userMessage map[string]string) (bool) {
	if len(userMessage) != 1 {
		return false
	}
	if _, itExists := userMessage["type"]; !itExists {
		return false
	}
	return true

}

func verifyDisconnectMessage(disconnectMessage map[string]string) (bool) {
	if len(disconnectMessage) != 1 {
		return false
	}
	if _, itExists := disconnectMessage["type"]; !itExists {
		return false
	}
	return true

}

func  verifyPublicMessage(publicMessage map[string]string) (bool) {
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

func verifyPrivateMessage(privateMessage map[string]string) (bool) {
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

func verifyNewRoomMessage(newRoomMessage map[string]string) (bool) {
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

func verifyInviteMessage(inviteMessage map[string]string) (bool) {
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

func verifyJoinMessage(joinMessage map[string]string) (bool) {
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

func verifyRoomUserMessage(roomUserMessage map[string]string) (bool) {
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

func verifyRoomMessage(roomMessage map[string]string) (bool) {
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

func verifyLeaveMessage(leaveMessage map[string]string) (bool) {
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
