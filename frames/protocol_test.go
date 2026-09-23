package frames

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

// TestRequestFrameValidation tests RequestFrame validation rules.
func TestRequestFrameValidation(t *testing.T) {
	tests := []struct {
		name        string
		frame       RequestFrame
		expectError bool
		errContains string
	}{
		{
			name: "valid GET request",
			frame: RequestFrame{
				Type:      FrameTypeRequest,
				RequestID: "req-001",
				Method:    "GET",
				Path:      "/api/v1/users",
			},
			expectError: false,
		},
		{
			name: "valid POST request with headers and body",
			frame: RequestFrame{
				Type:      FrameTypeRequest,
				RequestID: "req-002",
				Method:    "POST",
				Path:      "/api/v1/orders",
				Headers:   HeaderMap{"Content-Type": []string{"application/json"}},
				Body:      []byte(`{"item":"book"}`),
			},
			expectError: false,
		},
		{
			name: "valid lowercase method gets normalized",
			frame: RequestFrame{
				Type:      FrameTypeRequest,
				RequestID: "req-003",
				Method:    "delete",
				Path:      "/api/v1/items/42",
			},
			expectError: false,
		},
		{
			name: "invalid frame type",
			frame: RequestFrame{
				Type:      FrameTypeResponse,
				RequestID: "req-004",
				Method:    "GET",
				Path:      "/test",
			},
			expectError: true,
			errContains: "invalid frame type",
		},
		{
			name: "empty requestId",
			frame: RequestFrame{
				Type:      FrameTypeRequest,
				RequestID: "   ",
				Method:    "GET",
				Path:      "/test",
			},
			expectError: true,
			errContains: "requestId cannot be empty",
		},
		{
			name: "empty method",
			frame: RequestFrame{
				Type:      FrameTypeRequest,
				RequestID: "req-005",
				Method:    "",
				Path:      "/test",
			},
			expectError: true,
			errContains: "method cannot be empty",
		},
		{
			name: "unsupported HTTP method",
			frame: RequestFrame{
				Type:      FrameTypeRequest,
				RequestID: "req-006",
				Method:    "CONNECT",
				Path:      "/test",
			},
			expectError: true,
			errContains: "invalid HTTP method",
		},
		{
			name: "empty path",
			frame: RequestFrame{
				Type:      FrameTypeRequest,
				RequestID: "req-007",
				Method:    "GET",
				Path:      "   ",
			},
			expectError: true,
			errContains: "path cannot be empty",
		},
		{
			name: "path without leading slash",
			frame: RequestFrame{
				Type:      FrameTypeRequest,
				RequestID: "req-008",
				Method:    "GET",
				Path:      "api/v1/health",
			},
			expectError: true,
			errContains: "path must start with '/'",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.frame.Validate()
			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if tc.errContains != "" && !strings.Contains(err.Error(), tc.errContains) {
					t.Errorf("expected error to contain %q, got %q", tc.errContains, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}

// TestRequestFrameSerialization tests JSON marshaling and unmarshaling across various payloads.
func TestRequestFrameSerialization(t *testing.T) {
	largeBody := make([]byte, 64*1024) // 64 KB
	for i := range largeBody {
		largeBody[i] = byte('A' + (i % 26))
	}

	tests := []struct {
		name  string
		frame RequestFrame
	}{
		{
			name: "standard request with headers and body",
			frame: RequestFrame{
				Type:      FrameTypeRequest,
				RequestID: "req-test-12345",
				Method:    "GET",
				Path:      "/api/v1/health",
				Headers: HeaderMap{
					"User-Agent": []string{"Portalis-Test/1.0"},
				},
				Body: []byte(`{"status":"ok"}`),
			},
		},
		{
			name: "request with missing/nil headers",
			frame: RequestFrame{
				Type:      FrameTypeRequest,
				RequestID: "req-test-no-headers",
				Method:    "HEAD",
				Path:      "/ping",
				Headers:   nil,
				Body:      nil,
			},
		},
		{
			name: "request with empty binary body",
			frame: RequestFrame{
				Type:      FrameTypeRequest,
				RequestID: "req-test-empty-body",
				Method:    "POST",
				Path:      "/empty",
				Headers:   HeaderMap{"Content-Length": []string{"0"}},
				Body:      []byte{},
			},
		},
		{
			name: "request with very large payload (64KB)",
			frame: RequestFrame{
				Type:      FrameTypeRequest,
				RequestID: "req-test-large-body",
				Method:    "PUT",
				Path:      "/upload/blob",
				Headers:   HeaderMap{"Content-Type": []string{"application/octet-stream"}},
				Body:      largeBody,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.frame.Validate(); err != nil {
				t.Fatalf("expected valid frame: %v", err)
			}

			data, err := json.Marshal(tc.frame)
			if err != nil {
				t.Fatalf("failed to marshal frame: %v", err)
			}

			var decoded RequestFrame
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("failed to unmarshal frame: %v", err)
			}

			if decoded.RequestID != tc.frame.RequestID {
				t.Errorf("expected requestId %s, got %s", tc.frame.RequestID, decoded.RequestID)
			}
			if decoded.Method != tc.frame.Method {
				t.Errorf("expected method %s, got %s", tc.frame.Method, decoded.Method)
			}
			if decoded.Path != tc.frame.Path {
				t.Errorf("expected path %s, got %s", tc.frame.Path, decoded.Path)
			}
			if len(tc.frame.Body) > 0 && !bytes.Equal(decoded.Body, tc.frame.Body) {
				t.Errorf("decoded body does not match original body")
			}
		})
	}
}

// TestResponseFrameValidation tests status code boundaries and error conditions.
func TestResponseFrameValidation(t *testing.T) {
	tests := []struct {
		name        string
		frame       ResponseFrame
		expectError bool
		errContains string
	}{
		{
			name: "valid 200 OK response",
			frame: ResponseFrame{
				Type:       FrameTypeResponse,
				RequestID:  "req-resp-001",
				StatusCode: 200,
				Headers:    HeaderMap{"Content-Type": []string{"application/json"}},
				Body:       []byte(`{"result":"success"}`),
			},
			expectError: false,
		},
		{
			name: "valid lower boundary status code 100",
			frame: ResponseFrame{
				Type:       FrameTypeResponse,
				RequestID:  "req-resp-100",
				StatusCode: 100,
			},
			expectError: false,
		},
		{
			name: "valid upper boundary status code 599",
			frame: ResponseFrame{
				Type:       FrameTypeResponse,
				RequestID:  "req-resp-599",
				StatusCode: 599,
			},
			expectError: false,
		},
		{
			name: "invalid frame type",
			frame: ResponseFrame{
				Type:       FrameTypeRequest,
				RequestID:  "req-resp-bad-type",
				StatusCode: 200,
			},
			expectError: true,
			errContains: "invalid frame type",
		},
		{
			name: "empty requestId",
			frame: ResponseFrame{
				Type:       FrameTypeResponse,
				RequestID:  "   ",
				StatusCode: 200,
			},
			expectError: true,
			errContains: "requestId cannot be empty",
		},
		{
			name: "status code below 100 (99)",
			frame: ResponseFrame{
				Type:       FrameTypeResponse,
				RequestID:  "req-resp-under",
				StatusCode: 99,
			},
			expectError: true,
			errContains: "invalid HTTP statusCode",
		},
		{
			name: "status code zero (0)",
			frame: ResponseFrame{
				Type:       FrameTypeResponse,
				RequestID:  "req-resp-zero",
				StatusCode: 0,
			},
			expectError: true,
			errContains: "invalid HTTP statusCode",
		},
		{
			name: "status code above 599 (600)",
			frame: ResponseFrame{
				Type:       FrameTypeResponse,
				RequestID:  "req-resp-over",
				StatusCode: 600,
			},
			expectError: true,
			errContains: "invalid HTTP statusCode",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.frame.Validate()
			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if tc.errContains != "" && !strings.Contains(err.Error(), tc.errContains) {
					t.Errorf("expected error to contain %q, got %q", tc.errContains, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}

// TestResponseFrameSerialization tests JSON roundtrip for ResponseFrame.
func TestResponseFrameSerialization(t *testing.T) {
	resp := ResponseFrame{
		Type:       FrameTypeResponse,
		RequestID:  "req-test-resp-roundtrip",
		StatusCode: 404,
		Headers: HeaderMap{
			"Content-Type": []string{"text/plain; charset=utf-8"},
		},
		Body:  []byte("Not Found"),
		Error: "Resource does not exist",
	}

	data, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("failed to marshal ResponseFrame: %v", err)
	}

	var decoded ResponseFrame
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal ResponseFrame: %v", err)
	}

	if decoded.StatusCode != resp.StatusCode {
		t.Errorf("expected status %d, got %d", resp.StatusCode, decoded.StatusCode)
	}
	if decoded.Error != resp.Error {
		t.Errorf("expected error string %q, got %q", resp.Error, decoded.Error)
	}
	if !bytes.Equal(decoded.Body, resp.Body) {
		t.Errorf("decoded body does not match original body")
	}
}

// TestControlFrameValidation tests ControlFrame validation logic.
func TestControlFrameValidation(t *testing.T) {
	validPing := ControlFrame{Type: FrameTypePing}
	if err := validPing.Validate(); err != nil {
		t.Errorf("expected valid ping control frame, got: %v", err)
	}

	validPong := ControlFrame{Type: FrameTypePong}
	if err := validPong.Validate(); err != nil {
		t.Errorf("expected valid pong control frame, got: %v", err)
	}

	validError := ControlFrame{Type: FrameTypeError, Error: "tunnel disconnected"}
	if err := validError.Validate(); err != nil {
		t.Errorf("expected valid error control frame, got: %v", err)
	}

	invalidType := ControlFrame{Type: FrameTypeRequest}
	if err := invalidType.Validate(); err == nil {
		t.Errorf("expected error for invalid ControlFrame type, got nil")
	}
}

// TestFrameTypeHelpers tests String() and IsValid() helper methods.
func TestFrameTypeHelpers(t *testing.T) {
	types := []FrameType{
		FrameTypeRequest,
		FrameTypeResponse,
		FrameTypePing,
		FrameTypePong,
		FrameTypeError,
	}

	for _, ft := range types {
		if !ft.IsValid() {
			t.Errorf("expected FrameType %s to be valid", ft)
		}
		if ft.String() != string(ft) {
			t.Errorf("expected String() to return %s, got %s", string(ft), ft.String())
		}
	}

	invalid := FrameType("custom_unknown")
	if invalid.IsValid() {
		t.Errorf("expected custom_unknown to be invalid")
	}
}

// TestIsValidHTTPMethod tests the IsValidHTTPMethod helper.
func TestIsValidHTTPMethod(t *testing.T) {
	valid := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "get", "post"}
	for _, m := range valid {
		if !IsValidHTTPMethod(m) {
			t.Errorf("expected method %s to be valid", m)
		}
	}

	invalid := []string{"", "CONNECT", "TRACE", "FOO", "INVALID"}
	for _, m := range invalid {
		if IsValidHTTPMethod(m) {
			t.Errorf("expected method %s to be invalid", m)
		}
	}
}
