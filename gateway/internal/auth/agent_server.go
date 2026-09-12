package auth

import (
	"log"
	"net/http"
	"time"

	"github.com/dineTH2003-dev/Portalis/gateway/internal/config"
	"github.com/dineTH2003-dev/Portalis/gateway/internal/router"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  64 * 1024,
	WriteBufferSize: 64 * 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow cross-origin connection for CLI agents
	},
}

// AgentServer terminates agent WebSocket connections.
type AgentServer struct {
	cfg    *config.Config
	router *router.Router
}

// NewAgentServer creates a new AgentServer instance.
func NewAgentServer(cfg *config.Config, r *router.Router) *AgentServer {
	return &AgentServer{
		cfg:    cfg,
		router: r,
	}
}

// HandleWebSocket handles incoming agent WebSocket upgrade requests.
// TODO(contributor): [Issue #12]
//  1. Extract Bearer token from Authorization header
//  2. Parse and validate Tunnel Grant JWT using cfg.TunnelGrantSecret (HS256)
//  3. Extract claims: jti, sdn (subdomain), prt (local port)
//  4. Upgrade connection via upgrader.Upgrade(w, r, nil)
//  5. Register tunnel in s.router:
//     tunnel := &router.Tunnel{Subdomain: sdn, Conn: conn, LastHeartbeat: time.Now()}
//  6. Call Control Plane callback: POST /v1/internal/tunnel/sessions/:id/connected
//  7. Launch readLoop goroutine to process incoming response and pong frames
func (s *AgentServer) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}

	// Stub: Register dummy tunnel for demonstration until Issue #12 is implemented
	subdomain := r.URL.Query().Get("subdomain")
	if subdomain == "" {
		subdomain = "dev"
	}

	tunnel := &router.Tunnel{
		Subdomain:     subdomain,
		Conn:          conn,
		LastHeartbeat: time.Now(),
	}
	s.router.Register(subdomain, tunnel)
	log.Printf("Agent connected for subdomain: %s (stub)", subdomain)
}
