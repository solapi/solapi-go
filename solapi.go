package solapi

import (
	"github.com/solapi/solapi-go/cash"
	"github.com/solapi/solapi-go/messages"
	"github.com/solapi/solapi-go/storage"
)

// Client struct
type Client struct {
	Messages messages.Messages
	Storage  storage.Storage
	Cash     cash.Cash
}

// NewClient return a new client
func NewClient() *Client {
	client := Client{}
	return &client
}
