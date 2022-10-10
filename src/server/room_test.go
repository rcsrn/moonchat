package main

import (
	"testing"
	"strings"
	"fmt"
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

func TestVerifyRoomMember(t *testing.T) {
	if itExists := testRoom.verifyRoomMember("Person1"); itExists {
		t.Errorf("FAIL: the user does not exist in the room.")
	}
	
	testRoom.addRoomUser("Person1")
	
	if itExists := testRoom.verifyRoomMember("Person1"); !itExists {
		t.Errorf("FAIL: the user already exists in the room.")
	}
}

func TestGetMemberList(t *testing.T) {
	
}

func TestVerifyUserExistence(t *testing.T) {
	fmt.Println(testRoom.counter.users)
}

func TestAddRoomUser(t *testing.T) {
	
}


