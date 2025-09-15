package storages

import (
	"context"
	"fmt"

	"github.com/solapi/solapi-go/v2/internal/auth"
	"github.com/solapi/solapi-go/v2/internal/transport"
)

// Service provides storage-related APIs such as file upload.
type Service struct {
	baseURL string
	creds   auth.AuthenticationParameter
}

func NewService(baseURL string, creds auth.AuthenticationParameter) *Service {
	return &Service{baseURL: baseURL, creds: creds}
}

// Upload calls POST /storage/v1/files with JSON body.
func (s *Service) Upload(ctx context.Context, req UploadFileRequest) (UploadFileResponse, error) {
	url := fmt.Sprintf("%s/storage/v1/files", s.baseURL)
	httpReq := transport.DefaultRequest{URL: url, Method: "POST"}
	return transport.FetchJSON[UploadFileRequest, UploadFileResponse](ctx, s.creds, httpReq, &req)
}
