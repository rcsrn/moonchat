package main

import(
	"fmt"
	"testing"
	"github.com/rcsrn/moonchat/src/message"
	"strings"
)

//Test for checkidentify function
func TestCheckIdentify(t *testing.T) {
	var processor ServerProcessor = ServerProcessor{}
	
	gottenMessage := string(checkIdentify("Juan", &processor))
	mess := message.SuccesMessage{message.INFO_MESSAGE_TYPE, "Succes: username has been saved", message.IDENTIFY_MESSAGE_TYPE}
	rightMessage := string(mess.GetJSON())
	
	if value := strings.Compare(gottenMessage, rightMessage); value != 0 {
		t.Errorf("message that was gotten is %v and must be %v", gottenMessage, rightMessage)
	}
	
	gottenMessage2 := string(checkIdentify("Juan", &processor))
	mess2 := message.WarningMessageUsername{message.WARNING_MESSAGE_TYPE, "username already used", message.IDENTIFY_MESSAGE_TYPE, "Juan"}
	rightMessage2 := string(mess2.GetJSON())
	if value := strings.Compare(gottenMessage2, rightMessage2); value != 0 {
		t.Errorf("message that was gotten is %v and must be %v", gottenMessage2, rightMessage2)
	}	
}

//Test for adding a new user 
func TestAddUser(t *testing.T) {
	cleanUserMap()
	processor1 := ServerProcessor{}
	processor2 := ServerProcessor{}
	processor3 := ServerProcessor{}

	counter.users["Juan"] = &processor1
	counter.users["Brayan"] = &processor2
	counter.users["Pedro"] = &processor3

	if length := len(counter.users); length == 0 {
		t.Errorf("User has not been added to autentificated users")
	}
	if counter.users["Juan"] != &processor1 || counter.users["Brayan"] != &processor2 || counter.users["Pedro"] != &processor3 {
		t.Errorf("FAIL")
	}	
}

func TestGetUserProcessor(t *testing.T) {
	gottenProcessor, err := getUserProcessor("Juan")
	if err != nil || counter.users["Juan"] != gottenProcessor{
		t.Errorf("gotten processor is %v and it must be %v", gottenProcessor, counter.users["Juan"])
	}
}

func cleanUserMap() {
	for k := range counter.users {
		delete(counter.users, k)
	}
}

func cleanRoomMap () {
	for k, _ := range counter.rooms {
		delete(counter.users, k)
	}
}

func compareMessages(message1 []byte, message2 []byte) (int) {
	return strings.Compare(string(message1), string(message2))
}

func TestAddRoom(t *testing.T) {	
	initRooms()
	
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
	_ , err1:= createNewRoom("Juan", counter.users["Juan"], "")

	if err1 == nil {
		t.Errorf("FAIL: This should not happen.")
	}

	str1:= fmt.Sprintf("Succes: The room '%s'has been created succesfully.", "SALA100")
	succes := message.SuccesMessage{message.INFO_MESSAGE_TYPE, str1, message.NEW_ROOM_MESSAGE_TYPE}
	succesMessage := succes.GetJSON()
	
	gottenMessage1, err2 := createNewRoom("Juan", counter.users["Juan"], "SALA100")
	if err2 != nil {
		t.Errorf("FAIL: The room has not been created succesfully.")
	}
	
	if value := strings.Compare(string(succesMessage), string(gottenMessage1)); value != 0 {
		t.Errorf("FAIL: gotten name %v and it must be %v", string(gottenMessage1), string(succesMessage))
	}

	if length := len(counter.rooms); length == 0 {
		t.Errorf("FAIL")
	}
	
	found := false
	
	for nameRoom, _ := range counter.rooms {
		if value := strings.Compare(nameRoom, "SALA100"); value == 0 {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("FAIL: the room was not added succesfully to rooms.")
	}

	str2 := fmt.Sprintf("The '%s' already exists.", "SALA100")

	fail := message.RoomWarningMessage{message.WARNING_MESSAGE_TYPE, str2, message.NEW_ROOM_MESSAGE_TYPE, "SALA100"}

	failMessage := fail.GetJSON()
	
	gottenMessage2, err3 := createNewRoom("Juan", counter.users["Juan"], "SALA100")

	if err3 == nil {
		t.Errorf("FAIL: this should not happen.")
	}

	if value := strings.Compare(string(failMessage), string(gottenMessage2)); value != 0 {
		t.Errorf("FAIL: gotten name %v and it must be %v", string(gottenMessage2), string(failMessage))
	}
}

func TestJoinRoom (t *testing.T) {
	
}

func TestInviteUsersToRoom(t *testing.T) {
	
}

func TestVerifyRoomName(t *testing.T) {
	addRoom("roomTest", nil)
	if indicator := verifyRoomName("roomTest"); indicator {
		t.Errorf("FAIL")
	}
}

func TestVerifyStatus(t *testing.T) {
	if indicator := verifyStatus("SAD"); indicator {
		t.Error("FAIL")
	}
}
