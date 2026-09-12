package heartbeat

import (
	"log"
	"time"

	"github.com/dineTH2003-dev/Portalis/gateway/internal/router"
)

// Monitor tracks tunnel heartbeat pings and reaps stale sessions.
type Monitor struct {
	router *router.Router
	stop   chan struct{}
}

// NewMonitor initializes a new heartbeat monitor.
func NewMonitor(r *router.Router) *Monitor {
	return &Monitor{
		router: r,
		stop:   make(chan struct{}),
	}
}

// Start launches the 30-second heartbeat check loop in a background goroutine.
// TODO(contributor): [Issue #16]
//  1. Ticker fires every 30 seconds
//  2. Iterate over active tunnels
//  3. Send ControlFrame{Type: "ping"}
//  4. If now - tunnel.LastHeartbeat > 120s:
//     - Evict tunnel from router: m.router.Remove(subdomain)
//     - Close tunnel.Conn
//     - Notify Control Plane: POST /v1/internal/tunnel/sessions/:id/disconnected
func (m *Monitor) Start() {
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				log.Println("[Heartbeat] Running periodic health check...")
				// TODO(contributor): [Issue #16] Implement ping and stale session eviction
			case <-m.stop:
				return
			}
		}
	}()
}

// Stop terminates the heartbeat monitor.
func (m *Monitor) Stop() {
	close(m.stop)
}
