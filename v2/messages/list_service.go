package messages

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/solapi/solapi-go/v2/internal/transport"
)

// List calls GET /messages/v4/list with query parameters
func (s *Service) List(ctx context.Context, q ListQuery) (MessageListResponse, error) {
	params := url.Values{}
	if q.MessageID != "" {
		params.Set("messageId", q.MessageID)
	}
	if q.GroupID != "" {
		params.Set("groupId", q.GroupID)
	}
	if q.To != "" {
		params.Set("to", q.To)
	}
	if q.From != "" {
		params.Set("from", q.From)
	}
	if len(q.TypeIn) > 0 {
		params.Set("type", strings.Join(q.TypeIn, ","))
	}
	if q.DateType != "" {
		params.Set("dateType", q.DateType)
	}
	if q.StartDate != "" {
		params.Set("startDate", q.StartDate)
	}
	if q.EndDate != "" {
		params.Set("endDate", q.EndDate)
	}
	if q.StartKey != "" {
		params.Set("startKey", q.StartKey)
	}
	if q.Limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", q.Limit))
	}

	urlStr := fmt.Sprintf("%s/messages/v4/list", s.baseURL)
	if enc := params.Encode(); enc != "" {
		urlStr += "?" + enc
	}

	req := transport.DefaultRequest{URL: urlStr, Method: "GET"}
	return transport.FetchJSON[struct{}, MessageListResponse](ctx, s.creds, req, nil)
}
