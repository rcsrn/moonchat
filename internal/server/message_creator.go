package server

import "github.com/rcsrn/moonchat/internal/message"

type messageCreator struct {
}

func (creator *messageCreator) getRoomSuccesMessage(succes string, operation string, roomName string) ([]byte) {
	succesMessage := message.RoomSuccesMessage{message.INFO_TYPE, succes, operation, roomName}
	return succesMessage.GetJSON()
}

func (creator *messageCreator) getSuccesMessage(succes string, operation string) ([]byte) {
	succesMessage := message.SuccesMessage{message.INFO_TYPE, succes, operation}
	return succesMessage.GetJSON()
}

func (creator *messageCreator) getUsernameWarningMessage(warning string, operation string, userName string) ([]byte) {
	warningMessage := message.UsernameWarningMessage{message.WARNING_TYPE, warning, operation, userName}
	return warningMessage.GetJSON()
}

func (creator *messageCreator) getStatusWarningMessage(warning string, operation string, status string) ([]byte) {
	warningMessage := message.StatusWarningMessage{message.WARNING_TYPE, warning, operation, status}
	return warningMessage.GetJSON()
}

func (creator *messageCreator) getRoomWarningMessage(warning string, operation string, roomName string) ([]byte) {
	warningMessage := message.RoomWarningMessage{message.WARNING_TYPE, warning, operation, roomName}
	return warningMessage.GetJSON()
}

func (creator *messageCreator) getDisconnectedMessage(userName string) ([]byte) {
	disconnectedMessage := message.DisconnectedMessage{message.DISCONNECTED_TYPE, userName}
	return disconnectedMessage.GetJSON()
}

func (creator *messageCreator) getUserListMessage(typeOfList string, userList []string) ([]byte) {
	switch typeOfList {
	case message.ROOM_USER_LIST_TYPE:
		userListMessage := message.UserList{message.ROOM_USER_LIST_TYPE, userList}
		return userListMessage.GetJSON()
	default:
		userListMessage := message.UserList{message.USER_LIST_TYPE, userList}
		return userListMessage.GetJSON()
	}
}

func (creator *messageCreator) getPrivateMessage(private string, transmitter string) ([]byte) {
	privateMessage := message.NewMessage{message.PRIVATE_TYPE, transmitter, private}
	return privateMessage.GetJSON()
}

func (creator *messageCreator) getPublicMessage(userName string, public string) ([]byte) {
	publicMessage := message.NewMessage{message.PUBLIC_MESSAGE_FROM_TYPE, userName, public}
	return publicMessage.GetJSON()
}

func (creator *messageCreator) getNewStatusMessage(userName string, status string) ([]byte) {
	newStatusMessage := message.NewStatusMessage{message.NEW_STATUS_TYPE, userName, status}
	return newStatusMessage.GetJSON()
}

func (creator *messageCreator) getStatusErrorMessage(error string, operation string) ([]byte) {
	errorMessage := message.StatusErrorMessage{message.ERROR_TYPE, error, operation}
	return errorMessage.GetJSON()
}

func (creator *messageCreator) getRoomInvitationMessage(invitation string, host string, roomName string) ([]byte) {
	roomInvitationMessage := message.RoomInvitationMessage{message.INVITATION_TYPE, invitation, host, roomName}
	return roomInvitationMessage.GetJSON()
}

func (creator *messageCreator) getJoinedMessage(roomName string, userName string) ([]byte) {
	joinedMessage := message.JoinedRoomMessage{message.JOINED_ROOM_TYPE, roomName, userName}
	return joinedMessage.GetJSON()
}

func (creator *messageCreator) getRoomMessage(roomName string, userName string, messageToSend string) ([]byte) {
	roomMessage := message.RoomMessage{message.ROOM_MESSAGE_FROM_TYPE, roomName, userName, messageToSend}
	return roomMessage.GetJSON()
}

func (creator *messageCreator) getLeftRoomMessage(roomName string, userName string) ([]byte) {
	leftRoomMessage := message.LeftRoomMessage{message.LEFT_ROOM_TYPE, roomName, userName}
	return leftRoomMessage.GetJSON()
}

func (creator *messageCreator) getIdentifyErrorMessage() ([]byte) {
	errorMessage := message.IdentifyErrorMessage{message.ERROR_TYPE, "User not identified"}
	return errorMessage.GetJSON()
}
