package frames

import (
	"errors"
	"fmt"
	"strings"
)

// Validate checks whether a RequestFrame is well-formed.
// TODO(contributor): [Issue #2] Expand validation logic:
//   - Ensure r.Type == FrameTypeRequest
//   - Ensure r.RequestID is non-empty
//   - Ensure r.Method is a valid HTTP method (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS)
//   - Ensure r.Path is non-empty and starts with '/'
func (r *RequestFrame) Validate() error {
	if r.Type != FrameTypeRequest {
		return fmt.Errorf("invalid frame type for RequestFrame: expected %s, got %s", FrameTypeRequest, r.Type)
	}
	if strings.TrimSpace(r.RequestID) == "" {
		return errors.New("requestId cannot be empty")
	}
	if strings.TrimSpace(r.Method) == "" {
		return errors.New("method cannot be empty")
	}
	if strings.TrimSpace(r.Path) == "" {
		return errors.New("path cannot be empty")
	}
	return nil
}

// Validate checks whether a ResponseFrame is well-formed.
// TODO(contributor): [Issue #2] Expand validation logic:
//   - Ensure r.Type == FrameTypeResponse
//   - Ensure r.RequestID is non-empty
//   - Ensure r.StatusCode is within valid HTTP bounds (100 to 599)
func (r *ResponseFrame) Validate() error {
	if r.Type != FrameTypeResponse {
		return fmt.Errorf("invalid frame type for ResponseFrame: expected %s, got %s", FrameTypeResponse, r.Type)
	}
	if strings.TrimSpace(r.RequestID) == "" {
		return errors.New("requestId cannot be empty")
	}
	if r.StatusCode < 100 || r.StatusCode > 599 {
		return fmt.Errorf("invalid HTTP statusCode: %d", r.StatusCode)
	}
	return nil
}
