package client

import (
	"errors"
)

type Client struct {
	processor *ClientProcessor
	identified bool
}

func GetClientInstance() *Client {
	return &Client{&ClientProcessor{}, false}
}

func (client *Client) Connect() error {
	return errors.New("")
}

func (client *Client) ProcessMessage(message []string) {

}

func (client *Client) IsIdentified() bool {
	return client.identified
}

