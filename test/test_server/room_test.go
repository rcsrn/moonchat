package server

import (
	"testing"
	"strings"
	"github.com/rcsrn/moonchat/internal/server"
)

var testRoom *server.Room

func TestGetRoomInstance(t *testing.T) {
	testRoom = server.GetRoomInstance("SalaX")
	if testRoom == nil {
		t.Errorf("FAIL: nil has been gotten.")
	}
	if testRoom.GetRoomName() != "SalaX" {
		t.Errorf("FAIL: the names are different.")
	}
}

func TestAddRoomUser(t *testing.T) {
	if length := len(testRoom.GetMembers()); length != 0 {
		t.Errorf("FAIL: there are no users in room.")
	}	
	
	testRoom.AddUser("Person1")
	testRoom.AddUser("Person2")
	testRoom.AddUser("Person3")

	if !testRoom.VerifyRoomMember("Person1") {
		t.Errorf("FAIL: the user has not been added correctly.")
	}
	if !testRoom.VerifyRoomMember("Person2") {
		t.Errorf("FAIL: the user has not been added correctly.")
	}
	if !testRoom.VerifyRoomMember("Person3") {
		t.Errorf("FAIL: the user has not been added correctly.")
	}
	
	if length := len(testRoom.GetMembers()); length != 3 {
		t.Errorf("FAIL: there are three users in the room.")
	}
}

func TestVerifyRoomMember(t *testing.T) {
	if itExists := testRoom.VerifyRoomMember(""); itExists {
		t.Errorf("FAIL: this user does not exist in the room.")
	}
	if itExists := testRoom.VerifyRoomMember("Person1"); !itExists {
		t.Errorf("FAIL: this user does not exist in the room.")
	}
	if itExists := testRoom.VerifyRoomMember("Person2"); !itExists {
		t.Errorf("FAIL: this user does not exist in the room.")
	}
	if itExists := testRoom.VerifyRoomMember("Person3"); !itExists {
		t.Errorf("FAIL: this user does not exist in the room.")
	}
}


func TestGetMemberList(t *testing.T) {
	concurrentUserList := make([]string, 3)
	concurrentUserList[0] = "Person1"
	concurrentUserList[1] = "Person2"
	concurrentUserList[2] = "Person3"
	gottenUserList := testRoom.GetMemberList()
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


