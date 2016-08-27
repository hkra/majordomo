// Package apiserver provides a generic implementation of an API server.
package apiserver

import (
	"net"
	"strings"

	"crypto/tls"
	"log"
	"net/http"
)

// Config contains API server configuration.
type Config struct {
	Port        string
	BindAddress string
	UseTLS      bool
	Name        string
}

// APIServer groups configuration and components for the app API server.
type APIServer struct {
	config *Config
	mux    *http.ServeMux
	server *http.Server
}

// New creates a new APIServer instance.
func New(config *Config) *APIServer {
	s := &APIServer{
		config: config,
		mux:    http.NewServeMux(),
	}

	serverAddress := net.JoinHostPort(config.BindAddress, strings.Trim(config.Port, " \t"))

	s.server = &http.Server{
		Addr: serverAddress,
	}

	if config.UseTLS {
		s.server.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	return s
}

// Start runs the server and starts listening.
func (s *APIServer) Start() {
	log.Printf("Starting %s on port [%s]", s.config.Name, s.config.Port)
}
