package server

import "github.com/rcsrn/moonchat/cmd/message"

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
	return true
}

func verifyStatusMessage(StatusMessage map[string]string) (bool) {
	return true
}

func verifyUserMessage(UserMessage map[string]string) (bool) {
	return true
}

func verifyDisconnectMessage(disconnectMessage map[string]string) (bool) {
	return true
}

func  verifyPublicMessage(publicMessage map[string]string) (bool) {
	return true
}

func verifyPrivateMessage(privateMessage map[string]string) (bool) {
	return true
}

func verifyNewRoomMessage(newRoomMessage map[string]string) (bool) {
	return true
}

func verifyInviteMessage(inviteMessage map[string]string) (bool) {
	return true
}

func verifyJoinMessage(joinMessage map[string]string) (bool) {
	return true
}

func verifyRoomUserMessage(roomUserMessage map[string]string) (bool) {
	return true
}

func verifyRoomMessage(roomMessage map[string]string) (bool) {
	return true
}

func verifyLeaveMessage(leaveMessage map[string]string) (bool) {
	return true
}
