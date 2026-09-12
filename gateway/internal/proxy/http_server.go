package proxy

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dineTH2003-dev/Portalis/gateway/internal/config"
	"github.com/dineTH2003-dev/Portalis/gateway/internal/router"
)

// HTTPServer handles public HTTP ingress and proxies requests to tunnels.
type HTTPServer struct {
	cfg    *config.Config
	router *router.Router
}

// NewHTTPServer initializes the ingress HTTP server handler.
func NewHTTPServer(cfg *config.Config, r *router.Router) *HTTPServer {
	return &HTTPServer{
		cfg:    cfg,
		router: r,
	}
}

// ServeHTTP handles incoming public requests and dispatches them to active tunnels.
// TODO(contributor): [Issue #14 & #15]
//  1. Extract subdomain from r.Host (e.g. "myapp.localhost:8080" -> "myapp")
//  2. Fallback to X-Forwarded-Subdomain header
//  3. Lookup tunnel in s.router:
//     - If not found, return 404 "Tunnel Not Found" HTML
//  4. Convert http.Request to frames.RequestFrame:
//     - Generate unique RequestID (timestamp + random hex)
//     - Read request body bytes
//     - Copy headers into frames.HeaderMap
//  5. Register buffered response channel (chan *frames.ResponseFrame, cap 1)
//  6. Write JSON frame to tunnel.Conn under tunnel.ConnMu lock
//  7. Wait for response with 30s timeout:
//     - If timeout: return 504 Gateway Timeout
//     - If response received: copy headers, status code, and body to w
func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	subdomain := extractSubdomain(r.Host)
	if subdomain == "" {
		http.Error(w, "Invalid Host header: missing subdomain", http.StatusBadRequest)
		return
	}

	tunnel, exists := s.router.Lookup(subdomain)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "<html><body><h2>Portalis: Tunnel %q Not Found</h2><p>The tunnel may be offline or expired.</p></body></html>", subdomain)
		return
	}

	_ = tunnel
	// TODO(contributor): [Issue #15] Implement multiplexing logic here
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Portalis Gateway: Tunnel %q active (multiplexing stub)", subdomain)
}

func extractSubdomain(host string) string {
	parts := strings.Split(host, ":")
	hostname := parts[0]
	subParts := strings.Split(hostname, ".")
	if len(subParts) > 1 {
		return subParts[0]
	}
	return ""
}
