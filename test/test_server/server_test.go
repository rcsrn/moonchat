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
	processor1 := ServerProcessor{}
	processor2 := ServerProcessor{}
	processor3 := ServerProcessor{}
	
	testServer.GetUsers()["Juan"] = &processor1
	testServer.GetUsers()["Brayan"] = &processor2
	testServer.GetUsers()["Pedro"] = &processor3
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
	gottenProcessor, err := testServer.getUserProcessor("Juan")
	if err != nil || testServer.GetUsers()["Juan"] != gottenProcessor{
		t.Errorf("gotten processor is %v and it must be %v", gottenProcessor, testServer.GetUsers()["Juan"])
	}
}

func TestGetRoom(t *testing.T) {
	newRoom := room{}
	testServer.GetRooms()["Sala1"] = &newRoom
	gottenRoom, err := testServer.getRoom("Sala1")
	if err != nil || gottenRoom != &newRoom {
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
	newRoom1 := room{}
	newRoom2 := room{}
	newRoom3 := room {}
	
	testServer.addRoom("SALA1", &newRoom1)
	testServer.addRoom("SALA2", &newRoom2)
	testServer.addRoom("SALA3", &newRoom3)
	if length := len(testServer.GetRooms()); length == 0 {
		t.Errorf("The room has not been added succesfully")
	}
	if testServer.GetRooms()["SALA1"] != &newRoom1 || testServer.GetRooms()["SALA2"] != &newRoom2 || testServer.GetRooms()["SALA3"] != &newRoom3 {
		t.Errorf("FAIL: rooms have not been added correctly.")
	}
}

func TestCreateNewRoom(t *testing.T) {
	if error := testServer.createNewRoom("", nil, ""); error == nil {
		t.Errorf("FAIL: invalid room name")
	}
	
	testServer.addRoom("X", nil)
	
	if error := testServer.createNewRoom("", nil, "X"); error == nil {
		t.Errorf("FAIL: the room already exists.")
	}
	
	if error := testServer.createNewRoom("", nil, "Y"); error != nil {
		t.Errorf("FAIL: the room should be created.")
	}
	
}

func TestAddUserToRoom(t *testing.T) {
	cleanRoomMap()
}

func TestVerifyUserName(t *testing.T) {
	cleanUserMap()
	processor1 := ServerProcessor{}
	testServer.GetUsers()["Pedro"] = &processor1
	if available := testServer.verifyUserName("Pedro"); available {
		t.Errorf("The username is not available.")
	}
	if available := testServer.verifyUserName("pedro"); !available {
		t.Errorf("The username should be available.")
	}
}

func TestVerifyRoomName(t *testing.T) {
	testServer.addRoom("roomTest", nil)
	if indicator := testServer.verifyRoomName("roomTest"); indicator {
		t.Errorf("FAIL")
	}
}

func TestVerifyStatus(t *testing.T) {
	if isValid := testServer.verifyStatus("AWAY"); !isValid {
		t.Errorf("This status should be valid.")
	}
	if isValid := testServer.verifyStatus("BUSY"); !isValid {
		t.Errorf("This status should be valid.")
	}
	if isValid := testServer.verifyStatus("ACTIVE"); !isValid {
		t.Errorf("This status should be valid.")
	}
	if isValid := testServer.verifyStatus("SAD"); isValid {
		t.Error("FAIL: this is not a valid status")
	}
}

func TestVerifyIdentifiedUsers(t *testing.T) {
	cleanUserMap()
	fillUserList()
	usersToVerify := make([]string, 3)
	if theyAllExist, user := testServer.verifyIdentifiedUsers(usersToVerify); theyAllExist || strings.Compare(user, "") != 0 {
		t.Errorf("This should not happen.")
	}
	fillUserArray(usersToVerify)
	fillUserList()
	if theyAllExist, user := testServer.verifyIdentifiedUsers(usersToVerify); theyAllExist || strings.Compare(user, "") == 0 {
		t.Errorf("This should not happen.")
	}
	cleanUserMap()
	testServer.GetUsers()["Kimberly"] = nil
	testServer.GetUsers()["Jack"] = nil
	testServer.GetUsers()["Pepe"] = nil
	if theyAllExist, user := testServer.verifyIdentifiedUsers(usersToVerify); !theyAllExist || strings.Compare(user, "") != 0 {
		t.Errorf("This should not happen.")
	}
}


