package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

type DefaultRequest struct {
	URL    string
	Method string
}

// FetchJSON performs an HTTP request with Authorization header, retries on 503,
// maps 4xx to ApiError and 5xx to DefaultError, and decodes JSON into out.
func FetchJSON(ctx context.Context, params auth.AuthenticationParameter, req DefaultRequest, body any, out any) error {
	if req.URL == "" || req.Method == "" {
		return errors.New("invalid request")
	}

	authz, err := auth.BuildAuthorizationHeader(params)
	if err != nil {
		return err
	}

	const maxRetry = 3
	for attempt := 0; attempt <= maxRetry; attempt++ {
		var buf *bytes.Reader
		if body != nil {
			b, e := json.Marshal(body)
			if e != nil {
				return e
			}
			buf = bytes.NewReader(b)
		} else {
			buf = bytes.NewReader(nil)
		}

		httpReq, err := http.NewRequestWithContext(ctx, req.Method, req.URL, buf)
		if err != nil {
			return err
		}
		httpReq.Header.Set("Authorization", authz)
		httpReq.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(httpReq)
		if err != nil {
			if attempt < maxRetry {
				continue
			}
			return err
		}

		var retErr error
		func() {
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusServiceUnavailable {
				if attempt < maxRetry {
					_, _ = io.Copy(io.Discard, resp.Body)
					retErr = errors.New("retryable-503")
					return
				}
			}
			if resp.StatusCode >= 400 && resp.StatusCode < 500 {
				var er struct {
					ErrorCode    string `json:"errorCode"`
					ErrorMessage string `json:"errorMessage"`
				}
				if decErr := json.NewDecoder(resp.Body).Decode(&er); decErr != nil {
					er.ErrorCode = "ParseError"
					er.ErrorMessage = decErr.Error()
				}
				retErr = &ApiError{ErrorCode: er.ErrorCode, ErrorMessage: er.ErrorMessage, HTTPStatus: resp.StatusCode, URL: req.URL}
				return
			}
			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				b, _ := io.ReadAll(resp.Body)
				retErr = &DefaultError{ErrorCode: "UnknownError", ErrorMessage: string(b), Context: map[string]any{"status": resp.StatusCode, "url": req.URL}}
				return
			}
			if out == nil {
				retErr = nil
				return
			}
			retErr = json.NewDecoder(resp.Body).Decode(out)
		}()

		if retErr == nil {
			return nil
		}
		if retErr.Error() == "retryable-503" {
			continue
		}
		return retErr
	}
	return errors.New("unreachable")
}
