package transport

// DefaultRequest is a minimal HTTP request descriptor used by transport layer.
type DefaultRequest struct {
	URL    string
	Method string
}
