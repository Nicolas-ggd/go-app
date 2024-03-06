package main

// Server represents a server instance with its configuration.
type Server struct {
	config *Config // Configuration for the server.
}

// NewServer creates a new Server instance with the provided configuration.
func NewServer(config *Config) (*Server, error) {
	// Validate configuration or implement error handling if necessary
	return &Server{config: config}, nil
}

// Config holds configuration options for the server.
type Config struct {
	ListenerAddr string // Address on which the server listens for connections.
}

// WithListenerAddr creates a copy of the Config with the updated ListenerAddr.
func (c *Config) WithListenerAddr(addr string) *Config {
	// Consider cloning the Config to avoid mutating the original
	return &Config{ListenerAddr: addr}
}

// NewConfig creates a new Config instance with default values.
func NewConfig() *Config {
	return &Config{
		ListenerAddr: ":7000", // Default listener address.
	}
}
