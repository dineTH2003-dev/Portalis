package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dineTH2003-dev/Portalis/agent/internal/api"
	"github.com/dineTH2003-dev/Portalis/agent/internal/config"
	"github.com/dineTH2003-dev/Portalis/agent/internal/tunnel"
)

const version = "0.1.0"

func printUsage() {
	fmt.Println("Portalis CLI - Developer Connectivity Platform")
	fmt.Println("==============================================")
	fmt.Println("Usage: portalis <command> [arguments]")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  login <token>           Authenticate CLI with an agent token")
	fmt.Println("  http <port> [--subdomain name]  Expose a local HTTP port to the internet")
	fmt.Println("  config get              Display current CLI configuration")
	fmt.Println("  version                 Print CLI version")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "version":
		fmt.Printf("portalis CLI v%s\n", version)

	case "config":
		if len(os.Args) >= 3 && os.Args[2] == "get" {
			fmt.Printf("Server URL:  %s\n", cfg.ServerURL)
			fmt.Printf("User Email:  %s\n", cfg.UserEmail)
			fmt.Println("Token:       [saved]")
		} else {
			printUsage()
		}

	case "login":
		// TODO(contributor): [Issue #19] Implement login flow
		if len(os.Args) < 3 {
			fmt.Println("Usage: portalis login <agent-token>")
			os.Exit(1)
		}
		token := os.Args[2]
		client := api.NewClient(cfg.ServerURL)
		res, err := client.AgentLogin(token)
		if err != nil {
			fmt.Printf("Login failed: %v\n", err)
			os.Exit(1)
		}
		cfg.AgentToken = token
		cfg.UserEmail = res.Email
		if err := cfg.Save(); err != nil {
			fmt.Printf("Failed to save config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✅ Successfully logged in as %s!\n", res.Email)

	case "http":
		// TODO(contributor): [Issue #20] Implement port forwarding command
		if len(os.Args) < 3 {
			fmt.Println("Usage: portalis http <port> [--subdomain name]")
			os.Exit(1)
		}
		port, err := strconv.Atoi(os.Args[2])
		if err != nil || port < 1 || port > 65535 {
			fmt.Println("Error: Invalid port number (must be 1-65535)")
			os.Exit(1)
		}
		t := tunnel.NewClient("ws://localhost:9000/v1/tunnel/ws", cfg.AgentToken, port, "dev")
		if err := t.Start(); err != nil {
			fmt.Printf("Tunnel failed: %v\n", err)
			os.Exit(1)
		}

	default:
		printUsage()
		os.Exit(1)
	}
}
