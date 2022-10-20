package server

import (
	//"fmt"
	"testing"
	"strings"
	"github.com/rcsrn/moonchat/internal/server"
)

var testServer *server.Server


func TestGetServerInstance(t *testing.T) {
	testServer = server.GetServerInstance()
	if testServer.GetRooms() == nil || testServer.GetUsers() == nil {
		t.Errorf("FAIL: server does not init correctly")
	}
}

func cleanUserMap() {
	for k := range (testServer.GetUsers()) {
		delete(testServer.GetUsers(), k)
	}
}

func cleanRoomMap () {
	for k := range testServer.GetRooms() {
		delete(testServer.GetRooms(), k)
	}
}

func fillUserList() {
	processor1 := server.GetServerProcessorInstance(nil, nil)
	processor2 := server.GetServerProcessorInstance(nil, nil)
	processor3 := server.GetServerProcessorInstance(nil, nil)
	
	testServer.GetUsers()["Juan"] = processor1
	testServer.GetUsers()["Brayan"] = processor2
	testServer.GetUsers()["Pedro"] = processor3
}

func fillUserArray(users []string) {
	users[0] = "Kimberly"
	users[1] = "Pepe"
	users[2] = "Jack"
}

//Test for adding a new user 
func TestAddUser(t *testing.T) {
	cleanUserMap()
	fillUserList()	
}

func TestGetUserProcessor(t *testing.T) {
	gottenProcessor, err := testServer.GetUserProcessor("Juan")
	if err != nil || testServer.GetUsers()["Juan"] != gottenProcessor{
		t.Errorf("gotten processor is %v and it must be %v", gottenProcessor, testServer.GetUsers()["Juan"])
	}
}

func TestGetRoom(t *testing.T) {
	newRoom := server.GetRoomInstance("Sala1")
	testServer.GetRooms()["Sala1"] = newRoom
	gottenRoom, err := testServer.GetRoom("Sala1")
	if err != nil || gottenRoom != newRoom {
		t.Errorf("The gotten room is wrong")
	}
}


func compareAllStrings(lines []string, element string) (bool){
	for i := 0 ; i < len(lines); i++ {
		if value:= strings.Compare(lines[i], element); value == 0 {
			return true
		}
	}
	return false
}

func compareMessages(message1 []byte, message2 []byte) (int) {
	return strings.Compare(string(message1), string(message2))
}

func TestAddRoom(t *testing.T) {	
	newRoom1 := server.GetRoomInstance("SALA1")
	newRoom2 := server.GetRoomInstance("SALA1")
	newRoom3 := server.GetRoomInstance("SALA1")
	
	testServer.AddRoom("SALA1", newRoom1)
	testServer.AddRoom("SALA2", newRoom2)
	testServer.AddRoom("SALA3", newRoom3)
	if length := len(testServer.GetRooms()); length == 0 {
		t.Errorf("The room has not been added succesfully")
	}
	if testServer.GetRooms()["SALA1"] != newRoom1 || testServer.GetRooms()["SALA2"] != newRoom2 || testServer.GetRooms()["SALA3"] != newRoom3 {
		t.Errorf("FAIL: rooms have not been added correctly.")
	}
}

func TestCreateNewRoom(t *testing.T) {
	if error := testServer.CreateNewRoom("", nil, ""); error == nil {
		t.Errorf("FAIL: invalid room name")
	}
	
	testServer.AddRoom("X", nil)
	
	if error := testServer.CreateNewRoom("", nil, "X"); error == nil {
		t.Errorf("FAIL: the room already exists.")
	}
	
	if error := testServer.CreateNewRoom("", nil, "Y"); error != nil {
		t.Errorf("FAIL: the room should be created.")
	}
	
}

func TestAddUserToRoom(t *testing.T) {
	cleanRoomMap()
}

func TestVerifyUserName(t *testing.T) {
	cleanUserMap()
	processor1 := server.GetServerProcessorInstance(nil, nil)
	testServer.GetUsers()["Pedro"] = processor1
	if available := testServer.VerifyUserName("Pedro"); available {
		t.Errorf("The username is not available.")
	}
	if available := testServer.VerifyUserName("pedro"); !available {
		t.Errorf("The username should be available.")
	}
}

func TestVerifyRoomName(t *testing.T) {
	testServer.AddRoom("roomTest", nil)
	if indicator := testServer.VerifyRoomName("roomTest"); indicator {
		t.Errorf("FAIL")
	}
}

func TestVerifyStatus(t *testing.T) {
	if isValid := testServer.VerifyStatus("AWAY"); !isValid {
		t.Errorf("This status should be valid.")
	}
	if isValid := testServer.VerifyStatus("BUSY"); !isValid {
		t.Errorf("This status should be valid.")
	}
	if isValid := testServer.VerifyStatus("ACTIVE"); !isValid {
		t.Errorf("This status should be valid.")
	}
	if isValid := testServer.VerifyStatus("SAD"); isValid {
		t.Error("FAIL: this is not a valid status")
	}
}

func TestVerifyIdentifiedUsers(t *testing.T) {
	cleanUserMap()
	fillUserList()
	usersToVerify := make([]string, 3)
	if theyAllExist, user := testServer.VerifyIdentifiedUsers(usersToVerify); theyAllExist || strings.Compare(user, "") != 0 {
		t.Errorf("This should not happen.")
	}
	fillUserArray(usersToVerify)
	fillUserList()
	if theyAllExist, user := testServer.VerifyIdentifiedUsers(usersToVerify); theyAllExist || strings.Compare(user, "") == 0 {
		t.Errorf("This should not happen.")
	}
	cleanUserMap()
	testServer.GetUsers()["Kimberly"] = nil
	testServer.GetUsers()["Jack"] = nil
	testServer.GetUsers()["Pepe"] = nil
	if theyAllExist, user := testServer.VerifyIdentifiedUsers(usersToVerify); !theyAllExist || strings.Compare(user, "") != 0 {
		t.Errorf("This should not happen.")
	}
}


