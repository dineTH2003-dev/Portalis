package frames

import (
	"encoding/json"
	"testing"
)

// TestRequestFrameSerialization tests JSON marshaling and unmarshaling.
// TODO(contributor): [Issue #2] Add table-driven tests covering edge cases:
//   - Missing headers
//   - Empty binary bodies
//   - Very large body byte slices
func TestRequestFrameSerialization(t *testing.T) {
	req := RequestFrame{
		Type:      FrameTypeRequest,
		RequestID: "req-test-12345",
		Method:    "GET",
		Path:      "/api/v1/health",
		Headers: HeaderMap{
			"User-Agent": []string{"Portalis-Test/1.0"},
		},
		Body: []byte(`{"status":"ok"}`),
	}

	if err := req.Validate(); err != nil {
		t.Fatalf("expected valid request frame, got: %v", err)
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("failed to marshal frame: %v", err)
	}

	var decoded RequestFrame
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal frame: %v", err)
	}

	if decoded.RequestID != req.RequestID {
		t.Errorf("expected requestId %s, got %s", req.RequestID, decoded.RequestID)
	}
}

// TestResponseFrameValidation tests status code boundaries.
// TODO(contributor): [Issue #2] Add tests for:
//   - StatusCode < 100 (should fail)
//   - StatusCode > 599 (should fail)
//   - Empty RequestID (should fail)
func TestResponseFrameValidation(t *testing.T) {
	resp := ResponseFrame{
		Type:       FrameTypeResponse,
		RequestID:  "req-test-12345",
		StatusCode: 200,
	}

	if err := resp.Validate(); err != nil {
		t.Errorf("expected valid response frame, got: %v", err)
	}
}
