package main

import (
	"testing"
	"strings"
)

var testRoom room

func TestGetRoomInstance(t *testing.T) {
	testRoom := getRoomInstance("SalaX")
	if testRoom == nil {
		t.Errorf("FAIL: nil has been gotten.")
	}
	if indicator := strings.Compare(testRoom.roomName, "SalaX"); indicator != 0 {
		t.Errorf("FAIL: the names are different.")
	}
}

func TestInit(t *testing.T) {
	testRoom.init()
	if testRoom.counter.users == nil {
		t.Errorf("FAIL: nil has been gotten.")
	}
}

func TestAddRoomUser(t *testing.T) {
	if length := len(testRoom.counter.users); length != 0 {
		t.Errorf("FAIL: there are no users in room.")
	}
	testRoom.addRoomUser("Person1")
	testRoom.addRoomUser("Person2")
	
	if length := len(testRoom.counter.users); length != 2 {
		t.Errorf("FAIL: there are two users in the room.")
	}
}


func TestVerifyRoomMember(t *testing.T) {
	if itExists := testRoom.verifyRoomMember("Person3"); itExists {
		t.Errorf("FAIL: the user does not exist in the room.")
	}
	
	testRoom.addRoomUser("Person3")
	
	if itExists := testRoom.verifyRoomMember("Person3"); !itExists {
		t.Errorf("FAIL: the user already exists in the room.")
	}
}

func TestGetMemberList(t *testing.T) {
	
}

func TestVerifyUserExistence(t *testing.T) {
	if itExists := testRoom.verifyUserExistence(""); itExists {
		t.Errorf("FAIL: this user does not exist in the room.")
	}
	if itExists := testRoom.verifyUserExistence("Person1"); !itExists {
		t.Errorf("FAIL: this user does not exist in the room.")
	}
	if itExists := testRoom.verifyUserExistence("Person2"); !itExists {
		t.Errorf("FAIL: this user does not exist in the room.")
	}
	if itExists := testRoom.verifyUserExistence("Person3"); !itExists {
		t.Errorf("FAIL: this user does not exist in the room.")
	}
}


