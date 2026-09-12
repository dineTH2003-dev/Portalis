package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config stores persistent CLI preferences and credentials.
type Config struct {
	ServerURL  string `json:"serverUrl"`
	UserEmail  string `json:"userEmail,omitempty"`
	AgentToken string `json:"agentToken,omitempty"`
}

// DefaultPath returns the default config path: ~/.portalis/config.json
func DefaultPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".portalis", "config.json"), nil
}

// Load reads the configuration from disk.
// TODO(contributor): [Issue #18]
//   - If file does not exist, return default config with ServerURL = "http://localhost:4310"
//   - Read and unmarshal JSON
func Load() (*Config, error) {
	path, err := DefaultPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{ServerURL: "http://localhost:4310"}, nil
		}
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Save writes the configuration to ~/.portalis/config.json with safe file permissions (0600).
// TODO(contributor): [Issue #18]
//   - Ensure directory ~/.portalis exists with 0700 permissions
//   - Marshal config with indentation
//   - Write file with 0600 permissions
func (c *Config) Save() error {
	path, err := DefaultPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}
