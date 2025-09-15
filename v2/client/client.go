package client

import (
	"context"
	"errors"

	"github.com/solapi/solapi-go/v2/internal/auth"
	"github.com/solapi/solapi-go/v2/messages"
	"github.com/solapi/solapi-go/v2/storages"
)

const defaultBaseURL = "https://api.solapi.com"

// Client aggregates all solapi services under one entrypoint.
type Client struct {
	baseURL  string
	creds    auth.AuthenticationParameter
	Messages *messages.Service
	Storages *storages.Service
}

// NewClient initializes with default base URL.
func NewClient(apiKey, apiSecret string) *Client {
	return newClientWithBaseURL(defaultBaseURL, apiKey, apiSecret)
}

// newClientWithBaseURL is a test-only helper to override baseURL.
func newClientWithBaseURL(baseURL, apiKey, apiSecret string) *Client {
	creds := auth.AuthenticationParameter{ApiKey: apiKey, ApiSecret: apiSecret}
	c := &Client{baseURL: baseURL, creds: creds}
	c.Messages = messages.NewService(baseURL, creds)
	c.Storages = storages.NewService(baseURL, creds)
	return c
}

// Send is a convenience method delegating to Messages.Send with Background context.
func (c *Client) Send(input any, opts ...messages.SendOptions) (messages.DetailGroupMessageResponse, error) {
	// preflight validation: block empty recipient anywhere in ToList, and block when both To empty and ToList empty
	switch m := input.(type) {
	case messages.Message:
		if len(m.ToList) > 0 {
			for _, to := range m.ToList {
				if to == "" {
					return messages.DetailGroupMessageResponse{}, errors.New("recipient contains empty value")
				}
			}
		} else if m.To == "" {
			return messages.DetailGroupMessageResponse{}, errors.New("recipient is required")
		}
	case []messages.Message:
		for _, each := range m {
			if len(each.ToList) > 0 {
				for _, to := range each.ToList {
					if to == "" {
						return messages.DetailGroupMessageResponse{}, errors.New("recipient contains empty value")
					}
				}
			} else if each.To == "" {
				return messages.DetailGroupMessageResponse{}, errors.New("recipient is required")
			}
		}
	}
	return c.Messages.Send(context.Background(), input, opts...)
}

// List is a convenience method delegating to Messages.List with Background context.
func (c *Client) List(q messages.ListQuery) (messages.MessageListResponse, error) {
	return c.Messages.List(context.Background(), q)
}
