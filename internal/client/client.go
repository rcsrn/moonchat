package client

import (
	//	"errors"
	"net"
	"fmt"
)

type Client struct {
	processor *ClientProcessor
	identified bool
}

func GetClientInstance() *Client {
	return &Client{&ClientProcessor{}, false}
}

func (client *Client) Connect(host string) error {
	connection, error := net.Dial("tcp", host)
        if error != nil {
		fmt.Println(error.Error())
                return error
        }
	client.processor.setConnection(connection)
	return nil
}

func (client *Client) ProcessMessage(message []string) {

}

func (client *Client) IsIdentified() bool {
	return client.identified
}

