package transport

import (
	"fmt"
)

type ApiError struct {
	ErrorCode    string
	ErrorMessage string
	HTTPStatus   int
	URL          string
}

func (e *ApiError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("api error %s (%d) %s: %s", e.ErrorCode, e.HTTPStatus, e.URL, e.ErrorMessage)
}

type DefaultError struct {
	ErrorCode    string
	ErrorMessage string
	Context      map[string]any
}

func (e *DefaultError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.ErrorCode == "" {
		return e.ErrorMessage
	}
	return fmt.Sprintf("%s: %s", e.ErrorCode, e.ErrorMessage)
}
