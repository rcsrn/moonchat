package client

import (
	"net"
)

type Client struct {
	processor *ClientProcessor
	identified bool
}

func GetClientInstance() *Client {
	return &Client{&ClientProcessor{}, false}
}

func (client *Client) GetProcessor() *ClientProcessor {
	return client.processor
}

func (client *Client) Connect(host string) error {
	connection, error := net.Dial("tcp", host)
        if error != nil {
                return error
        }
	client.processor.setConnection(connection)
	return nil
}

func (client *Client) ProcessMessage(message []string) {
	client.processor.ProcessMessage(message)
}

func (client *Client) IsIdentified() bool {
	return client.identified
}

