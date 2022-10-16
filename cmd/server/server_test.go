package main

import(
	//	"fmt"
	"testing"
	"strings"
)

func initRooms() {
	counter.rooms = make(map[string]*room)
}

func cleanUserMap() {
	for k := range counter.users {
		delete(counter.users, k)
	}
}

func cleanRoomMap () {
	for k := range counter.rooms {
		delete(counter.rooms, k)
	}
}

func fillUserList() {
	processor1 := ServerProcessor{}
	processor2 := ServerProcessor{}
	processor3 := ServerProcessor{}
	
	counter.users["Juan"] = &processor1
	counter.users["Brayan"] = &processor2
	counter.users["Pedro"] = &processor3
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
	gottenProcessor, err := getUserProcessor("Juan")
	if err != nil || counter.users["Juan"] != gottenProcessor{
		t.Errorf("gotten processor is %v and it must be %v", gottenProcessor, counter.users["Juan"])
	}
}

func TestGetRoom(t *testing.T) {
	initRooms()
	newRoom := room{}
	counter.rooms["Sala1"] = &newRoom
	gottenRoom, err := getRoom("Sala1")
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
	
	addRoom("SALA1", &newRoom1)
	addRoom("SALA2", &newRoom2)
	addRoom("SALA3", &newRoom3)
	if length := len(counter.rooms); length == 0 {
		t.Errorf("The room has not been added succesfully")
	}
	if counter.rooms["SALA1"] != &newRoom1 || counter.rooms["SALA2"] != &newRoom2 || counter.rooms["SALA3"] != &newRoom3 {
		t.Errorf("FAIL")
	}
}

func TestCreateNewRoom(t *testing.T) {
	
}

func TestAddUserToRoom(t *testing.T) {
	cleanRoomMap()
	
	roomTest := room{}
	roomTest.init()
	userProcessor := ServerProcessor{}
	
	counter.rooms["roomTest"] = &roomTest
	
	error1 := addUserToRoom("Juan", "roomTest", &userProcessor)
	if error1 == nil {
		t.Errorf("FAIL: %v", error1.Error())
	}

	roomTest.counter.invitedUsers["Juan"] = &userProcessor
	
	error2 := addUserToRoom("Juan", "roomTest", &userProcessor)

	if error2 != nil {
		t.Errorf("FAIL: this error should not been thrown.")
	}

	if processor, itExists := roomTest.counter.users["Juan"]; !itExists || processor != &userProcessor {
		t.Errorf("FAIL: the user has not been added correctly.")
	}
	
	error3 := addUserToRoom("Juan", "", &userProcessor)
	
	if error3 == nil {
		t.Errorf("FAIL: %v", error3.Error())
	}
	
}

func TestVerifyUserName(t *testing.T) {
	cleanUserMap()
	processor1 := ServerProcessor{}
	counter.users["Pedro"] = &processor1
	if available := verifyUserName("Pedro"); available {
		t.Errorf("The username is not available.")
	}
	if available := verifyUserName("pedro"); !available {
		t.Errorf("The username should be available.")
	}
}

func TestVerifyRoomName(t *testing.T) {
	addRoom("roomTest", nil)
	if indicator := verifyRoomName("roomTest"); indicator {
		t.Errorf("FAIL")
	}
}

func TestVerifyStatus(t *testing.T) {
	if isValid := verifyStatus("AWAY"); !isValid {
		t.Errorf("This status should be valid.")
	}
	if isValid := verifyStatus("BUSY"); !isValid {
		t.Errorf("This status should be valid.")
	}
	if isValid := verifyStatus("ACTIVE"); !isValid {
		t.Errorf("This status should be valid.")
	}
	if isValid := verifyStatus("SAD"); isValid {
		t.Error("FAIL")
	}
}

func TestVerifyIdentifiedUsers(t *testing.T) {
	cleanUserMap()
	fillUserList()
	usersToVerify := make([]string, 3)
	if theyAllExist, user := verifyIdentifiedUsers(usersToVerify); theyAllExist || strings.Compare(user, "") != 0 {
		t.Errorf("This should not be happen.")
	}
	fillUserArray(usersToVerify)
	fillUserList()
	if theyAllExist, user := verifyIdentifiedUsers(usersToVerify); theyAllExist || strings.Compare(user, "") == 0 {
		t.Errorf("This should not be happen.")
	}
	cleanUserMap()
	counter.users["Kimberly"] = nil
	counter.users["Jack"] = nil
	counter.users["Pepe"] = nil
	if theyAllExist, user := verifyIdentifiedUsers(usersToVerify); !theyAllExist || strings.Compare(user, "") != 0 {
		t.Errorf("This should not be happen.")
	}
}


