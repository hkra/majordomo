// Package apiserver provides a generic implementation of an API server.
package apiserver

import (
	"log"
	"net/http"
)

// Config contains API server configuration.
type Config struct {
	Port uint
	Name string
}

// APIServer groups configuration and components for the app API server.
type APIServer struct {
	config *Config
	mux    *http.ServeMux
}

// New creates a new APIServer instance.
func New(config *Config) *APIServer {
	s := &APIServer{
		config: config,
		mux:    http.NewServeMux(),
	}
	return s
}

// Start runs the server and starts listening.
func (s *APIServer) Start() {
	log.Printf("Starting %s on port [%d]", s.config.Name, s.config.Port)
}
