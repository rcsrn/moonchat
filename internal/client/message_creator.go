package client

import "github.com/rcsrn/moonchat/internal/message"

type messageCreator struct {
}

func (creator *messageCreator) getDisconnectMessage() []byte {
	disconnectMessage := message.DisconnectMessage{message.DISCONNECT_TYPE}
	return disconnectMessage.GetJSON()
}

func (creator *messageCreator) getLeaveRoomMessage(roomname string) []byte {
	leaveRoomMessage := message.LeaveMessage{message.LEAVE_ROOM_TYPE, roomname}
	return leaveRoomMessage.GetJSON()
}

func (creator *messageCreator) getStatusMessage(status string) []byte {
	statusMessage := message.StatusMessage{message.STATUS_TYPE, status}
	return statusMessage.GetJSON()
}

func (creator *messageCreator) getUserListMessage() []byte {
	userListMessage := message.UserMessage{message.USERS_TYPE}
	return userListMessage.GetJSON()
}

func (creator *messageCreator) getPrivateMessage(receptor string, messageToSend string) []byte {
	privateMessage := message.PrivateMessage{message.MESSAGE_TYPE, receptor, messageToSend}
	return privateMessage.GetJSON()
}

func (creator *messageCreator) getNewRoomMessage(roomname string) []byte {
	newRoomMessage := message.NewRoomMessage{message.NEW_ROOM_TYPE, roomname}
	return newRoomMessage.GetJSON()
}

func (creator *messageCreator) getInvitationMessage(roomName string, userNames []string) []byte {
	invitationMessage := message.InviteToRoomMessage{message.INVITE_TYPE, roomName, userNames}
	return invitationMessage.GetJSON()
}

func (creator *messageCreator) getRoomMessage(roomName string, messageToSend string) []byte {
	roomMessage := message.ToRoomMessage{message.ROOM_MESSAGE_TYPE, roomName, messageToSend}
	return roomMessage.GetJSON()
}

func (creator *messageCreator) getPublicMessage(messageToSend string) []byte {
	publicMessage := message.PublicMessage{message.PUBLIC_TYPE, messageToSend}
	return publicMessage.GetJSON()
}

func (creator *messageCreator) getIdentifyMessage(userName string) []byte {
	identifyMessage := message.IdentifyMessage{message.IDENTIFY_TYPE, userName}
	return identifyMessage.GetJSON()
}



