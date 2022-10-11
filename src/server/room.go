package main

import (
	"sync"
	"fmt"
	//"strings"
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
	room.counter.blocker.RLock()
	if _, itExists := room.counter.users[userName]; itExists {
		return true
	}
	room.counter.blocker.RUnlock()
	return false
}

func (room *room) getMemberList() ([]string) {
	room.counter.blocker.RLock()
	memberList := make([]string, len(room.counter.users))
	i := 0
	for userName, _ := range(room.counter.users) {
		memberList[i] = userName
		i++
	}
	room.counter.blocker.RUnlock()
	return memberList
}

func (room *room) verifyInvitedUser(userName string) (bool) {
	if _, userHasBeenInvited := room.counter.invitedUsers[userName]; userHasBeenInvited {
		return true
	}
	return false
}

func (room *room) addUser(userName string) {
	userProcessor, _ := getUserProcessor(userName)
	fmt.Println("INTENTA TOMAR EL LOCK PERO NO PUEDE")
	room.counter.blocker.Lock()
	fmt.Println("INICIO DESBLOQUEO")
	room.counter.users[userName] = userProcessor
	room.counter.blocker.Unlock()
	fmt.Println("TERMINO DESBLOQUEO")
}
