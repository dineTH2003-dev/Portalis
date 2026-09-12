package router

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Tunnel represents an active agent tunnel connection.
type Tunnel struct {
	Subdomain     string
	Conn          *websocket.Conn
	ConnMu        sync.Mutex
	LastHeartbeat time.Time
	// PendingReqs maps RequestID to a response channel
	// TODO(contributor): [Issue #15] Add PendingReqs sync.Map or map[string]chan *frames.ResponseFrame
}

// Router provides thread-safe mapping between subdomains and active tunnels.
type Router struct {
	mu      sync.RWMutex
	tunnels map[string]*Tunnel
}

// NewRouter initializes an empty in-memory routing table.
func NewRouter() *Router {
	return &Router{
		tunnels: make(map[string]*Tunnel),
	}
}

// Register adds an active tunnel for the given subdomain.
// TODO(contributor): [Issue #13]
//   - Acquire write lock (r.mu.Lock())
//   - Check for existing subdomain collisions
//   - Store tunnel pointer in r.tunnels
func (r *Router) Register(subdomain string, t *Tunnel) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tunnels[subdomain] = t
}

// Lookup retrieves an active tunnel by subdomain.
// TODO(contributor): [Issue #13]
//   - Acquire read lock (r.mu.RLock())
//   - Return tunnel and boolean indicating presence
func (r *Router) Lookup(subdomain string) (*Tunnel, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.tunnels[subdomain]
	return t, ok
}

// Remove evicts a tunnel from the routing table.
// TODO(contributor): [Issue #13]
//   - Acquire write lock (r.mu.Lock())
//   - Delete subdomain key from r.tunnels
func (r *Router) Remove(subdomain string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.tunnels, subdomain)
}
