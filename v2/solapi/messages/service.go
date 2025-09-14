package messages

import (
	"context"
	"errors"
	"fmt"
	"runtime"

	"github.com/solapi/solapi-go/v2/internal/auth"
	"github.com/solapi/solapi-go/v2/internal/transport"
)

// Service handles message related APIs.
// It is intentionally small to satisfy v2 line constraints.

type Service struct {
	baseURL string
	creds   auth.AuthenticationParameter
}

func NewService(baseURL string, creds auth.AuthenticationParameter) *Service {
	return &Service{baseURL: baseURL, creds: creds}
}

// Send accepts Message, []Message, or SendRequest and normalizes to the API call.
func (s *Service) Send(ctx context.Context, input any, opts ...SendOptions) (DetailGroupMessageResponse, error) {
	switch v := input.(type) {
	case SendRequest:
		return s.SendManyDetail(ctx, v)
	case Message:
		var o SendOptions
		if len(opts) > 0 {
			o = opts[0]
		}
		return s.SendManyDetail(ctx, SendRequest{
			Messages:        []Message{v},
			AllowDuplicates: o.AllowDuplicates,
			ScheduledDate:   o.ScheduledDate,
			ShowMessageList: o.ShowMessageList,
			AppId:           o.AppId,
		})
	case []Message:
		var o SendOptions
		if len(opts) > 0 {
			o = opts[0]
		}
		return s.SendManyDetail(ctx, SendRequest{
			Messages:        v,
			AllowDuplicates: o.AllowDuplicates,
			ScheduledDate:   o.ScheduledDate,
			ShowMessageList: o.ShowMessageList,
			AppId:           o.AppId,
		})
	default:
		return DetailGroupMessageResponse{}, errors.New("unsupported input type for Send")
	}
}

// SendManyDetail converts SDK-facing SendRequest (with AppId) to API payload (with agent)
// and calls messages/v4/send-many/detail
func (s *Service) SendManyDetail(ctx context.Context, req SendRequest) (DetailGroupMessageResponse, error) {
	ag := &apiAgent{
		SDKVersion: "go/2.0.0",
		OSPlatform: runtime.GOOS + " | " + runtime.Version(),
	}
	if req.AppId != "" {
		ag.AppID = req.AppId
	}
	payload := apiSendRequest{
		Messages:        req.Messages,
		AllowDuplicates: req.AllowDuplicates,
		ScheduledDate:   req.ScheduledDate,
		ShowMessageList: req.ShowMessageList,
		Agent:           ag,
	}
	url := fmt.Sprintf("%s/messages/v4/send-many/detail", s.baseURL)
	httpReq := transport.DefaultRequest{URL: url, Method: "POST"}
	res, err := transport.FetchJSON[apiSendRequest, DetailGroupMessageResponse](ctx, s.creds, httpReq, &payload)
	if err != nil {
		return DetailGroupMessageResponse{}, err
	}
	c := res.GroupInfo.Count
	if len(res.FailedMessageList) > 0 && c.Total == c.RegisteredFailed {
		return DetailGroupMessageResponse{}, &MessageNotReceivedError{FailedMessageList: res.FailedMessageList, TotalCount: len(res.FailedMessageList)}
	}
	return res, nil
}
