package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client communicates with the Portalis Control Plane REST API.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient initializes a new API client.
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// AgentLoginResponse represents response from /v1/auth/agent-login.
type AgentLoginResponse struct {
	Success bool   `json:"success"`
	Email   string `json:"email"`
}

// AgentLogin validates the agent token against the Control Plane.
// TODO(contributor): [Issue #19]
//   - Make POST request to c.baseURL + "/v1/auth/agent-login"
//   - Pass Authorization: Bearer <agentToken>
//   - Parse response and return email
func (c *Client) AgentLogin(token string) (*AgentLoginResponse, error) {
	reqBody, _ := json.Marshal(map[string]string{"token": token})
	resp, err := c.httpClient.Post(c.baseURL+"/v1/auth/agent-login", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("login failed with HTTP status: %d", resp.StatusCode)
	}

	var res AgentLoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// TunnelSessionResponse represents response from POST /v1/tunnel/sessions.
type TunnelSessionResponse struct {
	SessionID  string `json:"sessionId"`
	GrantToken string `json:"grantToken"`
	PublicURL  string `json:"publicUrl"`
}

// RequestTunnelSession requests a short-lived Tunnel Grant JWT from the Control Plane.
// TODO(contributor): [Issue #20]
//   - Make POST to /v1/tunnel/sessions with token, port, and subdomain
//   - Return TunnelSessionResponse with GrantToken
func (c *Client) RequestTunnelSession(token string, port int, subdomain string) (*TunnelSessionResponse, error) {
	// Stub until Issue #20 is implemented
	return &TunnelSessionResponse{
		SessionID:  "session-dev-123",
		GrantToken: "dev-grant-token",
		PublicURL:  fmt.Sprintf("http://%s.localhost:8080", subdomain),
	}, nil
}
