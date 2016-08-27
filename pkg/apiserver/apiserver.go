// Package apiserver provides a generic implementation of an API server.
package apiserver

import (
	"net"
	"strings"

	"crypto/tls"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
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
	config           *Config
	mux              *http.ServeMux
	server           *http.Server
	restfulContainer *restful.Container
}

// New creates a new APIServer instance.
func New(config *Config) *APIServer {
	s := &APIServer{
		config:           config,
		mux:              http.NewServeMux(),
		restfulContainer: restful.NewContainer(),
	}

	s.restfulContainer.Router(restful.CurlyRouter{})

	// Setup CORS filter.
	s.restfulContainer.Filter(restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		CookiesAllowed: false,
		Container:      s.restfulContainer,
	}.Filter)

	serverAddress := net.JoinHostPort(config.BindAddress, strings.Trim(config.Port, " \t"))
	s.server = &http.Server{
		Addr:    serverAddress,
		Handler: s.restfulContainer,
	}

	if config.UseTLS {
		s.server.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	return s
}

// RegisterHandlers calls a callback used to add API route handlers to the restful container.
func (s *APIServer) RegisterHandlers(registrationCallback func(*restful.Container)) {
	registrationCallback(s.restfulContainer)
}

// Start runs the server and starts listening.
func (s *APIServer) Start() {
	log.Printf("Starting %s on port [%s]", s.config.Name, s.config.Port)

	if s.config.UseTLS {
		// TODO s.server.ListenAndServeTLS()
	} else {
		if s.server.ListenAndServe() != nil {
			log.Fatal("ListenAndServe failed during API server startup")
		}
	}
}
