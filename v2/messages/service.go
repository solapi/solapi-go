package messages

import (
	"context"
	"errors"
	"fmt"
	"runtime"

	"net/http"

	"github.com/solapi/solapi-go/v2/internal/auth"
	"github.com/solapi/solapi-go/v2/internal/transport"
)

type Service struct {
	baseURL    string
	creds      auth.AuthenticationParameter
	httpClient *http.Client
}

func NewService(baseURL string, creds auth.AuthenticationParameter) *Service {
	return &Service{baseURL: baseURL, creds: creds, httpClient: http.DefaultClient}
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

func (s *Service) SendManyDetail(ctx context.Context, req SendRequest) (DetailGroupMessageResponse, error) {
	// validate recipients: each message must have non-empty To or non-empty ToList
	for _, m := range req.Messages {
		if m.To == "" && len(m.ToList) == 0 {
			return DetailGroupMessageResponse{}, errors.New("recipient is required")
		}
		if len(m.ToList) > 0 {
			for _, to := range m.ToList {
				if to == "" {
					return DetailGroupMessageResponse{}, errors.New("recipient in ToList cannot be empty")
				}
			}
		}
	}

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
	ctx = transport.WithHTTPClient(ctx, s.httpClient)
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
