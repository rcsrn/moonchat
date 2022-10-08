package main

import (
	"sync"
)

type room struct {
	roomUsers mapCounter
	roomName string
}

type mapCounter struct {
	blocker sync.RWMutex
	users map[string]*ServerProcessor
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

func (room *room) addUser(userName string) {
	
}

