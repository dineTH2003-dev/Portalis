package tunnel

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Client manages the persistent WebSocket tunnel connection to the Gateway.
type Client struct {
	gatewayWSURL string
	grantToken   string
	localPort    int
	subdomain    string
	wsConn       *websocket.Conn
	httpClient   *http.Client
	stopChan     chan struct{}
}

// NewClient initializes a tunnel client instance.
func NewClient(gatewayWSURL, grantToken string, localPort int, subdomain string) *Client {
	return &Client{
		gatewayWSURL: gatewayWSURL,
		grantToken:   grantToken,
		localPort:    localPort,
		subdomain:    subdomain,
		httpClient: &http.Client{
			Timeout: 25 * time.Second,
		},
		stopChan: make(chan struct{}),
	}
}

// Start connects to the Gateway WebSocket and begins the frame forwarding loop.
// TODO(contributor): [Issue #20 & #22]
//   - Connect to c.gatewayWSURL with Header "Authorization: Bearer " + c.grantToken
//   - Implement exponential backoff reconnect loop if connection drops
//   - Launch c.readLoop() in a background goroutine
func (c *Client) Start() error {
	log.Printf("[Agent] Connecting tunnel for subdomain %q (target localhost:%d)...", c.subdomain, c.localPort)
	fmt.Println("🚀 Portalis Tunnel starting (stub)...")
	return nil
}

// readLoop reads frames from the Gateway and dispatches local HTTP requests.
// TODO(contributor): [Issue #21]
//  1. Loop: c.wsConn.ReadMessage()
//  2. Parse frame type:
//     - If "request": dispatch goroutine c.handleRequest(frame)
//     - If "ping": immediately answer c.wsConn.WriteJSON(ControlFrame{Type: "pong"})
//  3. In handleRequest:
//     - Construct http.Request to http://127.0.0.1:<localPort> + path
//     - Copy headers and body
//     - Execute request with c.httpClient
//     - If port down: send ResponseFrame with StatusCode: 502
//     - Package status, headers, and body into ResponseFrame
//     - Write ResponseFrame back to c.wsConn
func (c *Client) readLoop() {
	// Stub until Issue #21 is implemented
}

// Stop closes the tunnel connection cleanly.
func (c *Client) Stop() {
	close(c.stopChan)
	if c.wsConn != nil {
		c.wsConn.Close()
	}
}
