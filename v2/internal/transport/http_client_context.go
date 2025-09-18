package transport

import (
	"context"
	"net/http"
)

type ctxKey int

const httpClientKey ctxKey = 1

// WithHTTPClient stores *http.Client in context for transport to use.
func WithHTTPClient(ctx context.Context, c *http.Client) context.Context {
	if c == nil {
		return ctx
	}
	return context.WithValue(ctx, httpClientKey, c)
}

// httpClientFromContext retrieves *http.Client from context or returns http.DefaultClient.
func httpClientFromContext(ctx context.Context) *http.Client {
	if v := ctx.Value(httpClientKey); v != nil {
		if c, ok := v.(*http.Client); ok && c != nil {
			return c
		}
	}
	return http.DefaultClient
}
