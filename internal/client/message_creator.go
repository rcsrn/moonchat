package client

type messageCreator struct {
}

func (creator *messageCreator) getDisconnectMessage() []byte {
	return nil
}

func (creator *messageCreator) getLeaveRoomMessage(roomname string) []byte {
	return nil
}

func (creator *messageCreator) getStatusMessage(status string) []byte {
	return nil
}

func (creator *messageCreator) getUserListMessage() []byte {
	return nil
}

func (creator *messageCreator) getPrivateMessage(receptor string, messageTosend string) []byte {
	return nil
}

func (creator *messageCreator) getNewRoomMessage(roomname string) []byte {
	return nil
}

func (creator *messageCreator) getInvitationMessage(roomName string, userNames []string) []byte {
	return nil
}

func (creator *messageCreator) getRoomMessage(roomName string, message string) []byte {
	return nil
}

func (creator *messageCreator) getPublicMessage(message string) []byte {
	return nil
}

