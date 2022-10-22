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
	case LEAVE_ROOM: processor.leaveRoom()
		return nil
	case STATUS: processor.changeStatus()
		return nil
	case USER_LIST: processor.requestUserList()
		return nil
	case PRIVATE: processor.sendPrivateMessage()
		return nil
	case NEW_ROOM: processor.createNewRoom()
		return nil
	case INVITE: processor.inviteUsersToRoom()
		return nil
	case ROOM_MESSAGE: processor.sendRoomMessage()
		return nil
	default: processor.sendPublicMessage()
		return nil
	}
}

func (processor *ClientProcessor) disconnect() {
	disconnectMessage := processor.creator.getDisconnectMessage()
	processor.sendMessage(disconnectMessage)
}

func (processor *ClientProcessor) leaveRoom() {
	leaveRoomMessage := processor.creator.getLeaveRoomMessage()
	processor.sendMessage(leaveRoomMessage)
}
 
func (processor *ClientProcessor) changeStatus() {
	
}

func (processor *ClientProcessor) requestUserList() {
	
}

func (processor *ClientProcessor) sendPrivateMessage() {
	
}

func (processor *ClientProcessor) createNewRoom() {
	
}

func (processor *ClientProcessor) inviteUsersToRoom() {
	
}

func (processor *ClientProcessor) sendRoomMessage() {
	
}

func (processor *ClientProcessor) sendPublicMessage() {
	
}


func createError(errorMessage string) error {
	return errors.New(errorMessage)
}
