package main

import (
	"sync"
	//	"strings"
)

type room struct {
	counter mapCounter
	roomName string
}

type mapCounter struct {
	blocker sync.RWMutex
	users map[string]*ServerProcessor
}

func getRoomInstance(roomName string) *room {
	counter := mapCounter{}
	var roomInstance room = room{counter, roomName}
	return &roomInstance
}

func (room *room) init() {
	room.counter.users = make (map[string]*ServerProcessor)
}

func (room *room) verifyRoomMember(userName string) (bool) {
	return false 
}

func (room *room) getMemberList() ([]string) {
	return nil
}

func (room *room) itHasBeenInvited(userName string) (bool) {
	return true
}

func (room *room) availableUsername(userName string) (bool) {
	room.counter.blocker.RLock()
	if _, itExists := room.counter.users[userName]; itExists {
		return false
	}
	room.counter.blocker.RUnlock()
	return true
}

func (room *room) addUser(userName string) (error) {
	return nil
}
