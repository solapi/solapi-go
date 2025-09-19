package storages

import (
	"net/http"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

// NewServiceWithHTTPClient initializes Service with a custom *http.Client.
// If httpClient is nil, http.DefaultClient is used.
func NewServiceWithHTTPClient(baseURL string, creds auth.AuthenticationParameter, httpClient *http.Client) *Service {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Service{baseURL: baseURL, creds: creds, httpClient: httpClient}
}
