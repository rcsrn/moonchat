package server

import (
	"testing"
	"encoding/json"
	"github.com/rcsrn/moonchat/cmd/message"
	"github.com/rcsrn/moonchat/internal/server"
	"strings"
)

var testProcessor *server.ServerProcessor 

func TestGetServerProcessorInstance(t *testing.T) {
	testProcessor = server.GetServerProcessorInstance(nil, nil)
	if testProcessor == nil {
		t.Errorf("FAIL: server processor has not been created correctly.")
	}
}


func TestUnmarshalJSON(t *testing.T) {	
	var m message.SuccesMessage = message.SuccesMessage{message.INFO_TYPE, "this is test", message.IDENTIFY_TYPE}
	
	jsonMessage, err1 := json.Marshal(m)
	
	if err1 != nil {
		t.Errorf("This should not happen, the gotten error is: %v", err1)
	}
	
	gottenMessage, err2 := testProcessor.UnmarshalJSON(jsonMessage)

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

func TestToArrayOfUsers(t *testing.T) {
	stringSlice := make([]string, 3)
	fillUserArray(stringSlice)	
}

