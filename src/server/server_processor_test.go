package main

import (
	"testing"
	"encoding/json"
	"github.com/rcsrn/moonchat/src/message"
	"strings"
	//"net"
	//"fmt"
)

var testProcessor ServerProcessor;

func TestUnmarshalJSON(t *testing.T) {	
	testProcessor = ServerProcessor{}
	var m message.InfoMessage = message.InfoMessage{message.INFO_MESSAGE_TYPE, "this is test", message.IDENTIFY_MESSAGE_TYPE}
	
	jsonMessage, err1 := json.Marshal(m)
	
	if err1 != nil {
		t.Errorf("This should not happen, the error gotten is: %v", err1)
	}
	
	gottenMessage, err2 := testProcessor.unmarshalJSON(jsonMessage)

	if err2 != nil {
		t.Errorf("the error gotten is: %v", err2)
	}

	for key, element := range gottenMessage {
		if key == "type" {
			if value := strings.Compare(m.Type, element); value != 0 {
				t.Errorf("failed")
			}
		} else if key == "message" {
			if value := strings.Compare(m.Message, element); value != 0 {
				t.Errorf("failed")
			}
		} else if key == "operation" {
			if value := strings.Compare(m.Operation, element); value != 0 {
				t.Errorf("failed")
			}
		}
	}
}

// func TestSendMessages(t *testing.T) {
// 	conn, err := net.Dial("tcp", "localhost:1234")
// 	if err != nil {
// 		t.Errorf("could not connect to server")
// 	}
// 	testProcessor = ServerProcessor{conn}
// 	mess := message.InfoMessage{message.INFO_MESSAGE_TYPE, "this is a test", message.NEW_USER_TYPE}
// 	message := message.GetInfoMessageJSON(mess)
// 	testProcessor.sendMessage(message)
// }


