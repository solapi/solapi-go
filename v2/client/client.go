package client

import (
	"context"
	"net/http"

	"github.com/solapi/solapi-go/v2/groups"
	"github.com/solapi/solapi-go/v2/internal/auth"
	"github.com/solapi/solapi-go/v2/internal/transport"
	"github.com/solapi/solapi-go/v2/messages"
	"github.com/solapi/solapi-go/v2/storages"
)

const defaultBaseURL = "https://api.solapi.com"

// Client aggregates all solapi services under one entrypoint.
type Client struct {
	baseURL    string
	creds      auth.AuthenticationParameter
	httpClient *http.Client
	Messages   *messages.Service
	Storages   *storages.Service
	Groups     *groups.Service
}

// NewClient initializes with default base URL.
func NewClient(apiKey, apiSecret string) *Client {
	return newClientWithBaseURL(defaultBaseURL, apiKey, apiSecret)
}

// newClientWithBaseURL is a test-only helper to override baseURL.
func newClientWithBaseURL(baseURL, apiKey, apiSecret string) *Client {
	creds := auth.AuthenticationParameter{ApiKey: apiKey, ApiSecret: apiSecret}
	c := &Client{baseURL: baseURL, creds: creds, httpClient: http.DefaultClient}
	c.Messages = messages.NewService(baseURL, creds)
	c.Storages = storages.NewService(baseURL, creds)
	c.Groups = groups.NewService(baseURL, creds)
	return c
}

// WithHTTPClient returns a shallow copy of Client using the provided http.Client.
// If nil is passed, the receiver's client is kept.
func (c *Client) WithHTTPClient(hc *http.Client) *Client {
	if hc == nil {
		return c
	}
	nc := *c
	nc.httpClient = hc
	// Reinitialize services to ensure direct calls use the custom http.Client
	nc.Messages = messages.NewServiceWithHTTPClient(nc.baseURL, nc.creds, hc)
	nc.Storages = storages.NewServiceWithHTTPClient(nc.baseURL, nc.creds, hc)
	nc.Groups = groups.NewServiceWithHTTPClient(nc.baseURL, nc.creds, hc)
	return &nc
}

// Send is a convenience method delegating to Messages.Send with Background context.
func (c *Client) Send(input any, opts ...messages.SendOptions) (messages.DetailGroupMessageResponse, error) {
	ctx := context.Background()
	ctx = transport.WithHTTPClient(ctx, c.httpClient)
	return c.Messages.Send(ctx, input, opts...)
}

// List is a convenience method delegating to Messages.List with Background context.
func (c *Client) List(q messages.ListQuery) (messages.MessageListResponse, error) {
	ctx := context.Background()
	ctx = transport.WithHTTPClient(ctx, c.httpClient)
	return c.Messages.List(ctx, q)
}
