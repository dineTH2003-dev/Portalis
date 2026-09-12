package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dineTH2003-dev/Portalis/gateway/internal/auth"
	"github.com/dineTH2003-dev/Portalis/gateway/internal/config"
	"github.com/dineTH2003-dev/Portalis/gateway/internal/heartbeat"
	"github.com/dineTH2003-dev/Portalis/gateway/internal/proxy"
	"github.com/dineTH2003-dev/Portalis/gateway/internal/router"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("🚀 Portalis Ingress Gateway (Data Plane)")
	fmt.Println("==================================================")

	cfg := config.Load()
	r := router.NewRouter()

	// 1. Start Heartbeat Monitor
	monitor := heartbeat.NewMonitor(r)
	monitor.Start()
	defer monitor.Stop()

	// 2. Start Agent WebSocket Listener
	agentServer := auth.NewAgentServer(cfg, r)
	http.HandleFunc("/v1/tunnel/ws", agentServer.HandleWebSocket)

	go func() {
		log.Printf("[Gateway] Agent WebSocket listener on :%s/v1/tunnel/ws", cfg.WSPort)
		if err := http.ListenAndServe(":"+cfg.WSPort, nil); err != nil {
			log.Fatalf("WebSocket listener failed: %v", err)
		}
	}()

	// 3. Start Public HTTP Ingress Listener
	httpServer := proxy.NewHTTPServer(cfg, r)
	go func() {
		log.Printf("[Gateway] Public HTTP ingress on :%s", cfg.HTTPPort)
		if err := http.ListenAndServe(":"+cfg.HTTPPort, httpServer); err != nil {
			log.Fatalf("HTTP ingress listener failed: %v", err)
		}
	}()

	// 4. Handle OS Interrupts
	// TODO(contributor): [Issue #17] Implement graceful shutdown:
	//   - Catch SIGINT, SIGTERM
	//   - Stop accepting new connections
	//   - Wait up to 15s for pending requests to drain
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Println("[Gateway] Shutting down Portalis Gateway...")
}
