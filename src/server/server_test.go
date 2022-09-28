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
	cleanUsersMap()
	var username string = "Kimberly"
	var processor ServerProcessor = ServerProcessor{}
	fmt.Println(counter.users)
	addUser(username, &processor)	
	if length := len(counter.users); length == 0 {
		t.Errorf("User has not been added to autentificated users")
	}
	fmt.Printf("The user added is: %v\n", counter.users)
}

//Cleans the users map of the server to run more tests
func cleanUsersMap() {
	for k := range counter.users {
		delete(counter.users, k)
	}
}

func compareMessages(message1 []byte, message2 []byte) (int) {
	return strings.Compare(string(message1), string(message2))
}
