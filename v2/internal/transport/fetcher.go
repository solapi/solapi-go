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

var errRetryable = errors.New("retryable")

// FetchJSON performs an HTTP request with Authorization header, retries on 503,
// maps 4xx to ApiError and 5xx to DefaultError, and decodes JSON into TRes.
func FetchJSON[TReq any, TRes any](ctx context.Context, params auth.AuthenticationParameter, req DefaultRequest, body *TReq) (TRes, error) {
	return FetchJSONWithClient[TReq, TRes](ctx, httpClientFromContext(ctx), params, req, body)
}

// FetchJSONWithClient is like FetchJSON but uses the provided *http.Client.
// Library users can configure timeouts, transports, and middlewares via this client.
func FetchJSONWithClient[TReq any, TRes any](ctx context.Context, httpClient *http.Client, params auth.AuthenticationParameter, req DefaultRequest, body *TReq) (TRes, error) {
	var zero TRes
	if req.URL == "" || req.Method == "" {
		return zero, errors.New("invalid request")
	}

	authz, err := auth.BuildAuthorizationHeader(params)
	if err != nil {
		return zero, err
	}

	const maxRetry = 3
	for attempt := 0; attempt <= maxRetry; attempt++ {
		var buf *bytes.Reader
		if body != nil {
			b, e := json.Marshal(body)
			if e != nil {
				return zero, e
			}
			buf = bytes.NewReader(b)
		} else {
			buf = bytes.NewReader(nil)
		}

		httpReq, err := http.NewRequestWithContext(ctx, req.Method, req.URL, buf)
		if err != nil {
			return zero, err
		}
		httpReq.Header.Set("Authorization", authz)
		httpReq.Header.Set("Content-Type", "application/json")

		// Use provided client; caller is responsible for sensible defaults (e.g., timeouts)
		resp, err := httpClient.Do(httpReq)
		if err != nil {
			if attempt < maxRetry {
				continue
			}
			return zero, err
		}

		var retErr error
		var result TRes
		func() {
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusServiceUnavailable {
				if attempt < maxRetry {
					_, _ = io.Copy(io.Discard, resp.Body)
					retErr = errRetryable
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
				b, readErr := io.ReadAll(resp.Body)
				if readErr != nil {
					retErr = &DefaultError{ErrorCode: "UnknownError", ErrorMessage: readErr.Error(), Context: map[string]any{"status": resp.StatusCode, "url": req.URL}}
					return
				}
				retErr = &DefaultError{ErrorCode: "UnknownError", ErrorMessage: string(b), Context: map[string]any{"status": resp.StatusCode, "url": req.URL}}
				return
			}
			decErr := json.NewDecoder(resp.Body).Decode(&result)
			if decErr != nil && !errors.Is(decErr, io.EOF) {
				retErr = decErr
				return
			}
			retErr = nil
		}()

		if retErr == nil {
			return result, nil
		}
		if errors.Is(retErr, errRetryable) {
			continue
		}
		return zero, retErr
	}
	return zero, errors.New("unreachable")
}
