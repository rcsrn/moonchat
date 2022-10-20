package server

import (
	"testing"
	"strings"
	"github.com/rcsrn/moonchat/internal/server"
)

var testRoom *room 

func TestGetRoomInstance(t *testing.T) {
	testRoom = GetRoomInstance("SalaX")
	if testRoom == nil {
		t.Errorf("FAIL: nil has been gotten.")
	}
	if testRoom.roomName != "SalaX" {
		t.Errorf("FAIL: the names are different.")
	}
}

func TestAddRoomUser(t *testing.T) {
	if length := len(testRoom.users.elements); length != 0 {
		t.Errorf("FAIL: there are no users in room.")
	}	
	
	testRoom.addUser("Person1")
	testRoom.addUser("Person2")
	testRoom.addUser("Person3")

	if !testRoom.users.contains("Person1") {
		t.Errorf("FAIL: the user has not been added correctly.")
	}
	if !testRoom.users.contains("Person2") {
		t.Errorf("FAIL: the user has not been added correctly.")
	}
	if !testRoom.users.contains("Person3") {
		t.Errorf("FAIL: the user has not been added correctly.")
	}
	
	if length := len(testRoom.users.elements); length != 3 {
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
	isRight := equalSlices(concurrentUserList, gottenUserList)
	if !isRight {
		t.Errorf("FAIL: the gotten list is %v and it should be %v .",
			gottenUserList, concurrentUserList)
	}
}

//auxiliar function of TestGetMemberlist. Returns true if both slices equal.
//Implemented this way the order of the elements in the slices do not matter.
func equalSlices(slice1 []string, slice2 []string) (bool) {
	length1 := len(slice1)
	length2 := len(slice2)
	
	if length1 != length2 {
		return false
	}
	
	var isFound bool = true
	var i int
	var j int
	for i = 0; i < length1; i++ {
		for j = 0; j < length2; j++ {
			if val := strings.Compare(slice1[i], slice2[j]); val == 0 {
				isFound = true
			}
		}
		if !isFound {
			return false
		}
	}
	return true
}


