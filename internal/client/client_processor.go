package client

import (
	"net"
	"errors"
)

type ClientProcessor struct {
	connection net.Conn
	identified bool
	username string
	creator *messageCreator
	verifier *messageVerifier
}

func getClientInstance(connection net.Conn) *ClientProcessor {
	return &ClientProcessor{connection, false, "", &messageCreator{}, &messageVerifier{}}
}

func (processor *ClientProcessor) setConnection(connection net.Conn) {
	processor.connection = connection
}

func (processor *ClientProcessor) sendMessage(message []byte) {
	processor.connection.Write(message)
}

func (processor *ClientProcessor) ProcessMessage(message []string) error {
	if !processor.verifier.isValid(message) {
		return createError("invalid")
	}
	firstWord := message[0]
	switch firstWord {
	case CLOSE: processor.disconnect()
		return nil
	case LEAVE_ROOM: processor.leaveRoom(message[1])
		return nil
	case STATUS: processor.changeStatus(message[1])
		return nil
	case USER_LIST: processor.requestUserList()
		return nil
	case PRIVATE: processor.sendPrivateMessage(message[1], message[2])
		return nil
	case NEW_ROOM: processor.createNewRoom(message[1])
		return nil
	case INVITE:
		processor.inviteUsersToRoom(message[1], message[2:])
		return nil
	case ROOM_MESSAGE: processor.sendRoomMessage(message[1], message[2])
		return nil
	default: processor.sendPublicMessage(message[0])
		return nil
	}
}

func (processor *ClientProcessor) disconnect() {
	disconnectMessage := processor.creator.getDisconnectMessage()
	processor.sendMessage(disconnectMessage)
}

func (processor *ClientProcessor) leaveRoom(roomName string) {
	leaveRoomMessage := processor.creator.getLeaveRoomMessage(roomName)
	processor.sendMessage(leaveRoomMessage)
}
 
func (processor *ClientProcessor) changeStatus(status string) {
	newStatusMessage:= processor.creator.getStatusMessage(status)
	processor.sendMessage(newStatusMessage)
}

func (processor *ClientProcessor) requestUserList() {
	userListMessage := processor.creator.getUserListMessage()
	processor.sendMessage(userListMessage)
}

func (processor *ClientProcessor) sendPrivateMessage(receptor string, messageTosend string) {
	privateMessage := processor.creator.getPrivateMessage(receptor, messageTosend)
	processor.sendMessage(privateMessage)
}

func (processor *ClientProcessor) createNewRoom(roomname string) {
	newRoomMessage := processor.creator.getNewRoomMessage(roomname)
	processor.sendMessage(newRoomMessage)
}

func (processor *ClientProcessor) inviteUsersToRoom(roomName string, userNames []string) {
	invitationMessage := processor.creator.getInvitationMessage(roomName, userNames)
	processor.sendMessage(invitationMessage)
}

func (processor *ClientProcessor) sendRoomMessage(roomName string, message string) {
	roomMessage := processor.creator.getRoomMessage(roomName, message)
	processor.sendMessage(roomMessage)
}

func (processor *ClientProcessor) sendPublicMessage(message string) {
	publicMessage := processor.creator.getPublicMessage(message)
	processor.sendMessage(publicMessage)
}

func createError(errorMessage string) error {
	return errors.New(errorMessage)
}
