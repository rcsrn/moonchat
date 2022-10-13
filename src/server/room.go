package main

import (
	"sync"
	"strings"
)

type room struct {
	counter mapCounter
	roomName string
}

type mapCounter struct {
	blocker sync.RWMutex
	users map[string]*ServerProcessor
	invitedUsers map[string]*ServerProcessor
}

func getRoomInstance(roomName string) *room {
	counter := mapCounter{}
	var roomInstance room = room{counter, roomName}
	return &roomInstance
}

func (room *room) init() {
	room.counter.users = make(map[string]*ServerProcessor)
	room.counter.invitedUsers = make(map[string]*ServerProcessor)
}

func (room *room) verifyRoomMember(userName string) (bool) {
	if _, itExists := room.counter.users[userName]; itExists {
		return true
	}
	return false
}

func (room *room) getMemberList() ([]string) {
	memberList := make([]string, len(room.counter.users))
	i := 0
	for userName, _ := range(room.counter.users) {
		memberList[i] = userName
		i++
	}
	return memberList
}

func (room *room) verifyInvitedUser(userName string) (bool) {
	if _, userHasBeenInvited := room.counter.invitedUsers[userName]; userHasBeenInvited {
		return true
	}
	return false
}

func (room *room) addUser(userName string, userProcessor *ServerProcessor) {
	room.counter.users[userName] = userProcessor
}

func (room *room) addInvitedUser(userName string, userProcessor *ServerProcessor) {
	room.counter.invitedUsers[userName] = userProcessor
}

func (room *room) sendToAllUsers(transmitter string, message []byte) {
	for userName, userProcessor := range (room.counter.users) {
		if val := strings.Compare(userName, transmitter); val == 0 {
			continue
		}
		userProcessor.sendMessage(message)
	}
}
