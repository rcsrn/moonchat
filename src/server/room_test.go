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

	processor1 := ServerProcessor{}
	processor2 := ServerProcessor{}
	processor3 := ServerProcessor{}
	
	testRoom.addUser("Person1", &processor1)
	testRoom.addUser("Person2", &processor2)
	testRoom.addUser("Person3", &processor3)

	if processor := testRoom.counter.users["Person1"]; processor != &processor1 {
		t.Errorf("FAIL: the user has not been added correctly.")
	}
	if processor := testRoom.counter.users["Person1"]; processor != &processor1 {
		t.Errorf("FAIL: the user has not been added correctly.")
	}
	if processor := testRoom.counter.users["Person1"]; processor != &processor1 {
		t.Errorf("FAIL: the user has not been added correctly.")
	}
	
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
	concurrentUserList := make([]string, 3)
	concurrentUserList[0] = "Person1"
	concurrentUserList[1] = "Person2"
	concurrentUserList[2] = "Person3"
	gottenUserList := testRoom.getMemberList()
	isRight := compareSlices(concurrentUserList, gottenUserList)
	if !isRight {
		t.Errorf("FAIL: the gotten list is %v and it should be %v .",
			gottenUserList, concurrentUserList)
	}
}

//axuliar function of TestGetMemberlist. Returns true if both slices are equal.
func compareSlices(slice1 []string, slice2 []string) (bool) {
	length1 := len(slice1)
	length2 := len(slice2)
	
	if length1 != length2 {
		return false
	}
	
	for i := 0; i < length1; i++ {
		if indicator := strings.Compare(slice1[i], slice2[i]); indicator != 0 {
			return false
		}
	}
	return true
}


