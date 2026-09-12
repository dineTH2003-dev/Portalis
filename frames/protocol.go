package frames

// FrameType specifies the role of a protocol frame in the Portalis wire format.
type FrameType string

const (
	FrameTypeRequest  FrameType = "request"
	FrameTypeResponse FrameType = "response"
	FrameTypePing     FrameType = "ping"
	FrameTypePong     FrameType = "pong"
	FrameTypeError    FrameType = "error"
)

// HeaderMap represents HTTP header keys mapped to string slices.
type HeaderMap map[string][]string

// RequestFrame transports an inbound HTTP request from the Gateway to the Agent.
type RequestFrame struct {
	Type      FrameType `json:"type"`
	RequestID string    `json:"requestId"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Headers   HeaderMap `json:"headers"`
	Body      []byte    `json:"body,omitempty"`
}

// ResponseFrame transports an HTTP response from the Agent back to the Gateway.
type ResponseFrame struct {
	Type       FrameType `json:"type"`
	RequestID  string    `json:"requestId"`
	StatusCode int       `json:"statusCode"`
	Headers    HeaderMap `json:"headers"`
	Body       []byte    `json:"body,omitempty"`
	Error      string    `json:"error,omitempty"`
}

// ControlFrame handles connection keepalive and control signaling.
type ControlFrame struct {
	Type  FrameType `json:"type"`
	Error string    `json:"error,omitempty"`
}

// ErrorFrame transports a terminal error condition between endpoints.
type ErrorFrame struct {
	Type      FrameType `json:"type"`
	RequestID string    `json:"requestId,omitempty"`
	Message   string    `json:"message"`
	Code      string    `json:"code,omitempty"`
}

// String returns the string representation of the frame type.
func (f FrameType) String() string {
	return string(f)
}

// IsValid checks if the frame type is one of the supported types.
func (f FrameType) IsValid() bool {
	switch f {
	case FrameTypeRequest, FrameTypeResponse, FrameTypePing, FrameTypePong, FrameTypeError:
		return true
	default:
		return false
	}
}
