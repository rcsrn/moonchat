package main

type messageVerifier struct {
	message map[string]string
}

func (verifier *messageVerifier) isValidMessage() (bool) {
	switch verifier.message {
	case message.IDENTIFY_TYPE:
		return verifyIdentifyMessage(message)
	case message.STATUS_TYPE:
		return verifyStatusMessage(message)
	case message.USERS_TYPE:
	
	case message.DISCONNECT_TYPE:
	
	case message.PUBLIC_TYPE:
	
	case message.MESSAGE_TYPE:
	
	case message.NEW_ROOM_TYPE:
	
	case message.INVITE_TYPE:
	
	case message.JOIN_ROOM_TYPE:
		
	case message.ROOM_USERS_TYPE:
		
	case message.ROOM_MESSAGE_TYPE:
	
	case message.LEAVE_ROOM_TYPE:
		
	default: disconnectClientCase(processor)	
	}
} 

func verifyIdentifyMessage(identifyMessage map[string]string) (bool) {
	
}

func verifyStatusMessage(StatusMessage map[string]string) (bool) {
	return nil
}

func verifyUserMessage(UserMessage map[string]string) (bool) {
	return nil
}

func verifyDisconnectMessage(disconnectMessage map[string]string) (bool) {
	return nil
}

func  verifyPublicMessage(publicMessage map[string]string) (bool) {
	return nil
}

func verifyNewRoomMessage(newRoomMessage map[string]string) (bool) {
	return nil
}

func verifyJoinMessage(joinMessage map[string]string) (bool) {
	return nil
}

func verifyRoomUserMessage(roomUserMessage map[string]string) (bool) {
	return nil
}

func verifyRoomMessage(roomMessage map[string]string) (bool) {
	return nil
}

func verifyLeaveMessage(leaveMessage map[string]string) (bool) {
	return nil
}
