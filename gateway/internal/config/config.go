package config

import (
	"os"
)

// Config holds operational parameters for the Portalis Gateway.
type Config struct {
	HTTPPort              string
	WSPort                string
	ServerAPIURL          string
	InternalGatewaySecret string
	TunnelGrantSecret     string
}

// Load reads Gateway parameters from environment variables with safe defaults.
// TODO(contributor): [Issue #11]
//   - Support CLI flags override (--http-port, --ws-port)
//   - Validate required secrets (fail fast if TunnelGrantSecret is missing in production)
func Load() *Config {
	return &Config{
		HTTPPort:              getEnv("PORTALIS_GATEWAY_HTTP_PORT", "8080"),
		WSPort:                getEnv("PORTALIS_GATEWAY_WS_PORT", "9000"),
		ServerAPIURL:          getEnv("PORTALIS_SERVER_API_URL", "http://localhost:4310"),
		InternalGatewaySecret: getEnv("INTERNAL_GATEWAY_SECRET", "dev_secret_gateway_change_in_prod"),
		TunnelGrantSecret:     getEnv("TUNNEL_GRANT_SECRET", "dev_secret_grant_change_in_prod"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
