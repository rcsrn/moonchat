package main

import(
	//"fmt"
	"testing"
	"github.com/rcsrn/moonchat/src/message"
	"strings"
)

//Test for checkidentify function
func TestCheckIdentify(t *testing.T) {
	var processor ServerProcessor = ServerProcessor{}
	gottenMessage := string(checkIdentify("Juan", &processor))
	mess := message.InfoMessage{message.INFO_MESSAGE_TYPE, "Succes: username has been saved", message.IDENTIFY_MESSAGE_TYPE}
	rightMessage := string(message.GetInfoMessageJSON(mess))

	if value := strings.Compare(gottenMessage, rightMessage); value != 0 {
		t.Errorf("message that was gotten is %v and must be %v", gottenMessage, rightMessage)
	}
	
	gottenMessage2 := string(checkIdentify("Juan", &processor))
	mess2 := message.WarningMessage{message.WARNING_MESSAGE_TYPE, "username already used", message.IDENTIFY_MESSAGE_TYPE, "Juan"}
	rightMessage2 := string(message.GetWarningMessageJSON(mess2))
	if value := strings.Compare(gottenMessage2, rightMessage2); value != 0 {
		t.Errorf("message that was gotten is %v and must be %v", gottenMessage2, rightMessage2)
	}
}

//func TestAddUser(t *testing.T) {
// 	var username string = "Username"
// 	var processor ServerProcessor = ServerProcessor{}
// 	addUser(username, &processor)
// 	if 
// }


