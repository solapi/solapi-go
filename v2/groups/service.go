package groups

import (
	"context"
	"fmt"
	"net/url"
	"runtime"

	"net/http"

	"github.com/solapi/solapi-go/v2/internal/auth"
	"github.com/solapi/solapi-go/v2/internal/transport"
	"github.com/solapi/solapi-go/v2/messages"
)

// Service exposes group-related endpoints
type Service struct {
	baseURL    string
	creds      auth.AuthenticationParameter
	httpClient *http.Client
}

func NewService(baseURL string, creds auth.AuthenticationParameter) *Service {
	return &Service{baseURL: baseURL, creds: creds, httpClient: http.DefaultClient}
}

type CreateGroupOptions struct {
	AllowDuplicates bool
	AppId           string
	CustomFields    map[string]string
}

// Create POST /messages/v4/groups
func (s *Service) Create(ctx context.Context, opt CreateGroupOptions) (CreateGroupResponse, error) {
	type body struct {
		SDKVersion      string            `json:"sdkVersion"`
		OSPlatform      string            `json:"osPlatform"`
		AllowDuplicates bool              `json:"allowDuplicates"`
		AppId           string            `json:"appId,omitempty"`
		CustomFields    map[string]string `json:"customFields,omitempty"`
	}
	b := body{
		SDKVersion:      "go/2.0.0",
		OSPlatform:      runtime.GOOS + " | " + runtime.Version(),
		AllowDuplicates: opt.AllowDuplicates,
		AppId:           opt.AppId,
		CustomFields:    opt.CustomFields,
	}
	urlStr := fmt.Sprintf("%s/messages/v4/groups", s.baseURL)
	req := transport.DefaultRequest{URL: urlStr, Method: "POST"}
	ctx = transport.WithHTTPClient(ctx, s.httpClient)
	return transport.FetchJSON[body, CreateGroupResponse](ctx, s.creds, req, &b)
}

// AddMessages PUT /messages/v4/groups/{groupId}/messages
func (s *Service) AddMessages(ctx context.Context, groupId string, reqBody AddGroupMessagesRequest) (GroupActionResponse, error) {
	urlStr := fmt.Sprintf("%s/messages/v4/groups/%s/messages", s.baseURL, groupId)
	req := transport.DefaultRequest{URL: urlStr, Method: "PUT"}
	ctx = transport.WithHTTPClient(ctx, s.httpClient)
	return transport.FetchJSON[AddGroupMessagesRequest, GroupActionResponse](ctx, s.creds, req, &reqBody)
}

// ListMessages GET /messages/v4/groups/{groupId}/messages
func (s *Service) ListMessages(ctx context.Context, groupId string, q ListMessagesQuery) (messages.MessageListResponse, error) {
	values := url.Values{}
	if q.StartKey != "" {
		values.Set("startKey", q.StartKey)
	}
	if q.Limit > 0 {
		values.Set("limit", fmt.Sprintf("%d", q.Limit))
	}
	urlStr := fmt.Sprintf("%s/messages/v4/groups/%s/messages", s.baseURL, groupId)
	if enc := values.Encode(); enc != "" {
		urlStr += "?" + enc
	}
	req := transport.DefaultRequest{URL: urlStr, Method: "GET"}
	ctx = transport.WithHTTPClient(ctx, s.httpClient)
	return transport.FetchJSON[struct{}, messages.MessageListResponse](ctx, s.creds, req, nil)
}

// Send POST /messages/v4/groups/{groupId}/send
func (s *Service) Send(ctx context.Context, groupId string) (messages.DetailGroupMessageResponse, error) {
	urlStr := fmt.Sprintf("%s/messages/v4/groups/%s/send", s.baseURL, groupId)
	req := transport.DefaultRequest{URL: urlStr, Method: "POST"}
	ctx = transport.WithHTTPClient(ctx, s.httpClient)
	return transport.FetchJSON[struct{}, messages.DetailGroupMessageResponse](ctx, s.creds, req, nil)
}

// Reserve POST /messages/v4/groups/{groupId}/schedule
func (s *Service) Reserve(ctx context.Context, groupId string, scheduledDate string) (messages.DetailGroupMessageResponse, error) {
	urlStr := fmt.Sprintf("%s/messages/v4/groups/%s/schedule", s.baseURL, groupId)
	req := transport.DefaultRequest{URL: urlStr, Method: "POST"}
	body := ScheduleRequest{ScheduledDate: scheduledDate}
	ctx = transport.WithHTTPClient(ctx, s.httpClient)
	return transport.FetchJSON[ScheduleRequest, messages.DetailGroupMessageResponse](ctx, s.creds, req, &body)
}

// CancelReservation DELETE /messages/v4/groups/{groupId}/schedule
func (s *Service) CancelReservation(ctx context.Context, groupId string) (messages.DetailGroupMessageResponse, error) {
	urlStr := fmt.Sprintf("%s/messages/v4/groups/%s/schedule", s.baseURL, groupId)
	req := transport.DefaultRequest{URL: urlStr, Method: "DELETE"}
	ctx = transport.WithHTTPClient(ctx, s.httpClient)
	return transport.FetchJSON[struct{}, messages.DetailGroupMessageResponse](ctx, s.creds, req, nil)
}

// ListGroups GET /messages/v4/groups
func (s *Service) ListGroups(ctx context.Context, q ListGroupsQuery) (ListGroupsResponse, error) {
	values := url.Values{}
	if q.StartKey != "" {
		values.Set("startKey", q.StartKey)
	}
	if q.Limit > 0 {
		values.Set("limit", fmt.Sprintf("%d", q.Limit))
	}
	urlStr := fmt.Sprintf("%s/messages/v4/groups", s.baseURL)
	if enc := values.Encode(); enc != "" {
		urlStr += "?" + enc
	}
	req := transport.DefaultRequest{URL: urlStr, Method: "GET"}
	ctx = transport.WithHTTPClient(ctx, s.httpClient)
	return transport.FetchJSON[struct{}, ListGroupsResponse](ctx, s.creds, req, nil)
}

// GetGroup GET /messages/v4/groups/{groupId}
func (s *Service) GetGroup(ctx context.Context, groupId string) (GroupResponse, error) {
	urlStr := fmt.Sprintf("%s/messages/v4/groups/%s", s.baseURL, groupId)
	req := transport.DefaultRequest{URL: urlStr, Method: "GET"}
	ctx = transport.WithHTTPClient(ctx, s.httpClient)
	return transport.FetchJSON[struct{}, GroupResponse](ctx, s.creds, req, nil)
}

// RemoveMessages DELETE /messages/v4/groups/{groupId}/messages
func (s *Service) RemoveMessages(ctx context.Context, groupId string, messageIds []string) (GroupActionResponse, error) {
	urlStr := fmt.Sprintf("%s/messages/v4/groups/%s/messages", s.baseURL, groupId)
	req := transport.DefaultRequest{URL: urlStr, Method: "DELETE"}
	body := RemoveGroupMessagesRequest{MessageIDs: messageIds}
	ctx = transport.WithHTTPClient(ctx, s.httpClient)
	return transport.FetchJSON[RemoveGroupMessagesRequest, GroupActionResponse](ctx, s.creds, req, &body)
}

// RemoveGroup DELETE /messages/v4/groups/{groupId}
func (s *Service) RemoveGroup(ctx context.Context, groupId string) (GroupResponse, error) {
	urlStr := fmt.Sprintf("%s/messages/v4/groups/%s", s.baseURL, groupId)
	req := transport.DefaultRequest{URL: urlStr, Method: "DELETE"}
	ctx = transport.WithHTTPClient(ctx, s.httpClient)
	return transport.FetchJSON[struct{}, GroupResponse](ctx, s.creds, req, nil)
}
