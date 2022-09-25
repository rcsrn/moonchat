package main

import(
	"fmt"
	"testing"
	"github.com/rcsrn/moonchat/src/message"
	"strings"
	"net"
	//"encoding/json"
	//"io"
)

//Test for checkidentify function
func TestCheckIdentify(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		t.Errorf("could not connect to server")
	}
	defer conn.Close()
	var processor ServerProcessor = ServerProcessor{conn}
	
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

//Test for adding a new user 
func TestAddUser(t *testing.T) {
	cleanUsersMap()
	var username string = "Username"
	var processor ServerProcessor = ServerProcessor{}
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

//Test for sending a message to all identified users
// func TestToAllUsers(t *testing.T) {
// 	cleanUsersMap()
// 	conn, err := net.Dial("tcp", "localhost:1234")
// 	if err != nil {
// 		t.Errorf("could not connect to server")
// 	}
// 	defer conn.Close()
// 	process := ServerProcessor{conn}
// 	addUser("Rodrigo", &process)
// 	mess := message.NewUserMessage{message.NEW_USER_TYPE, "Juan"}
// 	message1, err := json.Marshal(mess)
// 	if err != nil {
// 		t.Errorf("Message was not created succesfully")
// 	}
	
// 	//toAllUsers(message1)
// 	el := counter.users["Rodrigo"]
// 	fmt.Println(el)
// 	el.connection.Write(message1)

	
// 	fmt.Println("The length of identified user map is: ", len(counter.users))
// 	fmt.Println(counter.users)
// 	for user, processor := range counter.users {
// 		buffer := make([]byte, 1024)
// 		_ , err := processor.connection.Read(buffer)
		
// 		if err != nil {
// 			if err == io.EOF {
// 				t.Errorf("It is not possible to read from connection")				
// 			}
// 		}
// 		var message2 []byte
// 		err1 := json.Unmarshal(buffer, &message2)
// 		if err1 != nil {
// 			t.Errorf("Json was not unmarshaled succesfully")
// 		}
// 		if value := compareMessages(message1, message2); value != 0 {
// 			t.Errorf("The message recieved for the user %v is %v and must be %v", user, message2, message1)
// 		}
// 	}
// }

func compareMessages(message1 []byte, message2 []byte) (int) {
	return strings.Compare(string(message1), string(message2))
}
