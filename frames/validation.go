package frames

import (
	"errors"
	"fmt"
	"strings"
)

// ValidHTTPMethods defines the supported HTTP methods for RequestFrame.
var ValidHTTPMethods = map[string]struct{}{
	"GET":     {},
	"POST":    {},
	"PUT":     {},
	"DELETE":  {},
	"PATCH":   {},
	"HEAD":    {},
	"OPTIONS": {},
}

// IsValidHTTPMethod checks whether the provided string is an accepted HTTP method.
func IsValidHTTPMethod(method string) bool {
	_, exists := ValidHTTPMethods[strings.ToUpper(strings.TrimSpace(method))]
	return exists
}

// Validate checks whether a RequestFrame is well-formed.
// Ensures:
//   - r.Type == FrameTypeRequest
//   - r.RequestID is non-empty
//   - r.Method is a valid HTTP method (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS)
//   - r.Path is non-empty and starts with '/'
func (r *RequestFrame) Validate() error {
	if r.Type != FrameTypeRequest {
		return fmt.Errorf("invalid frame type for RequestFrame: expected %s, got %s", FrameTypeRequest, r.Type)
	}
	if strings.TrimSpace(r.RequestID) == "" {
		return errors.New("requestId cannot be empty")
	}
	method := strings.ToUpper(strings.TrimSpace(r.Method))
	if method == "" {
		return errors.New("method cannot be empty")
	}
	if !IsValidHTTPMethod(method) {
		return fmt.Errorf("invalid HTTP method: %s", r.Method)
	}
	path := strings.TrimSpace(r.Path)
	if path == "" {
		return errors.New("path cannot be empty")
	}
	if !strings.HasPrefix(path, "/") {
		return fmt.Errorf("path must start with '/': got %s", r.Path)
	}
	return nil
}

// Validate checks whether a ResponseFrame is well-formed.
// Ensures:
//   - r.Type == FrameTypeResponse
//   - r.RequestID is non-empty
//   - r.StatusCode is within valid HTTP bounds (100 to 599)
func (r *ResponseFrame) Validate() error {
	if r.Type != FrameTypeResponse {
		return fmt.Errorf("invalid frame type for ResponseFrame: expected %s, got %s", FrameTypeResponse, r.Type)
	}
	if strings.TrimSpace(r.RequestID) == "" {
		return errors.New("requestId cannot be empty")
	}
	if r.StatusCode < 100 || r.StatusCode > 599 {
		return fmt.Errorf("invalid HTTP statusCode: %d (must be between 100 and 599)", r.StatusCode)
	}
	return nil
}

// Validate checks whether a ControlFrame is well-formed.
func (c *ControlFrame) Validate() error {
	if c.Type != FrameTypePing && c.Type != FrameTypePong && c.Type != FrameTypeError {
		return fmt.Errorf("invalid frame type for ControlFrame: expected ping, pong, or error, got %s", c.Type)
	}
	return nil
}
