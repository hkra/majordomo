// Package app contains the broker server application.
package app

import (
	"log"

	"github.com/hkra/majordomo/cmd/broker/app/options"
)

// APIServer groups configuration and components for the app API server.
type APIServer struct {
	Port uint
}

// Config configures an APIServer using the values from an Options object.
func (s *APIServer) Config(options *options.Options) {
	s.Port = options.Port
}

// Start runs the server and starts listening.
func (s *APIServer) Start() {
	log.Printf("Starting majordomo command broker on port [%d]", s.Port)
}
