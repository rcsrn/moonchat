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
	if testRoom.counter.invitedUsers == nil {
		t.Errorf("FAIL: nil has been gotten.")
	}
}

func TestAddRoomUser(t *testing.T) {
	if length := len(testRoom.counter.users); length != 0 {
		t.Errorf("FAIL: there are no users in room.")
	}
	testRoom.addRoomUser("Person1")
	testRoom.addRoomUser("Person2")
	testRoom.addRoomUser("Person3")
	
	if length := len(testRoom.counter.users); length != 3 {
		t.Errorf("FAIL: there are three users in the room.")
	}
}

func TestVerifyRoomMember(t *testing.T) {
	if itExists := testRoom.verifyRoomMember(""); itExists {
		t.Errorf("FAIL: this user does not exist in the room.")
	}
	if itExists := testRoom.verifyRoomMember("Person1"); !itExists {
		t.Errorf("FAIL: this user does not exist in the room.")
	}
	if itExists := testRoom.verifyRoomMember("Person2"); !itExists {
		t.Errorf("FAIL: this user does not exist in the room.")
	}
	if itExists := testRoom.verifyRoomMember("Person3"); !itExists {
		t.Errorf("FAIL: this user does not exist in the room.")
	}
}


func TestGetMemberList(t *testing.T) {
	
}



