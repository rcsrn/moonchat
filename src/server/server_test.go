package main

import(
	"fmt"
	"testing"
	"github.com/rcsrn/moonchat/src/message"
	"strings"
	//"encoding/json"
	//reflect"
)

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

func fillUserList() {
	processor1 := ServerProcessor{}
	processor2 := ServerProcessor{}
	processor3 := ServerProcessor{}
	
	counter.users["Juan"] = &processor1
	counter.users["Brayan"] = &processor2
	counter.users["Pedro"] = &processor3
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
	_ , err1:= createNewRoom("Juan", counter.users["Juan"], "")

	if err1 == nil {
		t.Errorf("FAIL: Invalid room name.")
	}

	str1:= fmt.Sprintf("Succes: The room '%s'has been created succesfully.", "SALA100")
	succes := message.SuccesMessage{message.INFO_TYPE, str1, message.NEW_ROOM_TYPE}
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

	fail := message.RoomWarningMessage{message.WARNING_TYPE, str2, message.NEW_ROOM_TYPE, "SALA100"}

	failMessage := fail.GetJSON()
	
	gottenMessage2, err3 := createNewRoom("Juan", counter.users["Juan"], "SALA100")

	if err3 == nil {
		t.Errorf("FAIL: this should not happen.")
	}

	if value := strings.Compare(string(failMessage), string(gottenMessage2)); value != 0 {
		t.Errorf("FAIL: gotten name %v and it must be %v", string(gottenMessage2), string(failMessage))
	}
}

func TestJoinRoom(t *testing.T) {
	if error := joinRoom("", ""); error == nil {
		t.Errorf("FAIL: both invalid name of room and user.")
	}
	if error := joinRoom("", "Juan"); error == nil {
		t.Errorf("FAIL: invalid room name.")
	}
	if error := joinRoom("SALA1", "Jack"); error == nil {
		t.Errorf("FAIL: user is not in identified list.")
	}
	if error := joinRoom("SALA1", "Juan"); error == nil {
		t.Errorf("FAIL: the user has not been invited to room.")
	}
	//FALTA EL CASO CUANDO SI ESTA INVITADO
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

func fillUserArray(users []string) {
	users[0] = "Kimberly"
	users[1] = "Pepe"
	users[2] = "Jack"
}
 
