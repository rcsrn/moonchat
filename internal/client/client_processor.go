package client

import (
	"net"
	"errors"
	"encoding/json"
	"strings"
	"github.com/rcsrn/moonchat/internal/message"
	"os"
	"fmt"
)

type ClientProcessor struct {
	connection net.Conn
	creator *messageCreator
	verifier *messageVerifier
}

func getClientInstance(connection net.Conn) *ClientProcessor {
	return &ClientProcessor{connection, &messageCreator{}, &messageVerifier{}}
}

func (processor *ClientProcessor) ReadFromServer() {
	for {
		buffer := make([]byte, 1024)
		numBitsRead, error1 := processor.connection.Read(buffer)
		
		if error1 != nil {
			break
		}
		
		//It is necessary to specify the exact number of bits read by Read() method to unmarshal.
		messageReceived, error2 := processor.UnmarshalJSON(buffer[:numBitsRead])
		
		if error2 != nil {
			break
		}

		processor.throwServerMessage(messageReceived)
		
	}
}

func (processor *ClientProcessor) UnmarshalJSON(j []byte) (map[string]string, error) {
	var mapString map[string]string 
	error1 := json.Unmarshal(j, &mapString)
	if error1 != nil {
		if val := strings.Compare(error1.Error(), "json: cannot unmarshal array into Go value of type string"); val == 0 {
			var message message.InviteToRoomMessage
			error2 := json.Unmarshal(j, &message)
			if error2 != nil {
				return nil, error2
			}
			return convertMessageToMapString(message), nil
		} else {
			return nil, error1
		}
	}
	return mapString, nil
}

//auxiliar function to deal the case when it is necessary to work with an array of strings.
func convertMessageToMapString(message message.InviteToRoomMessage) (map[string]string) {
	mapString := make(map[string]string)
	mapString["type"] = message.Type
	mapString["roomname"] = message.Roomname
	mapString["usernames"] = strings.Join(message.Usernames, ",")
	return mapString
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
	case IDENTIFY: processor.identify(message[1])
		return nil
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
	case INVITE: processor.inviteUsersToRoom(message[1], message[2:])
		return nil
	case ROOM_MESSAGE: processor.sendRoomMessage(message[1], message[2])
		return nil
	default: processor.sendPublicMessage(message[0])
		return nil
	}
}

func (processor *ClientProcessor) throwServerMessage(messageReceived map[string]string) map[string]string {
	fmt.Println(messageReceived)
	return messageReceived
}

func (processor *ClientProcessor) identify(userName string) {
	identifyMessage := processor.creator.getIdentifyMessage(userName)
	processor.sendMessage(identifyMessage)
}

func (processor *ClientProcessor) disconnect() {
	disconnectMessage := processor.creator.getDisconnectMessage()
	processor.sendMessage(disconnectMessage)
	processor.connection.Close()
	os.Exit(0)
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
