// Package httpserver implements HTTP server.
package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"perx/internal/config"
	"perx/internal/service"
	"perx/internal/transport/http/api"
)

type Server struct {
	http   *http.Server
	config config.Config
}

// New returns a new instance of Server.
func New(cfg config.Config, services *service.Services) *Server {
	s := &Server{
		config: cfg,
	}

	// version 1
	r := api.NewAPI(cfg, services.Item)

	s.http = &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	return s
}

// Open validates the server options and begins listening on the bind address.
func (s *Server) Open() error {
	go func() {
		address := net.JoinHostPort(s.config.IP, s.config.HTTPPort)
		log.Printf("Start HTTP on %s\n", address)

		if err := s.http.ListenAndServe(); err != nil {
			log.Fatal("failed to http serve")
		}
	}()

	return nil
}

// Close gracefully shuts down the server.
func (s *Server) Close(ctx context.Context) error {
	log.Println("gracefully shuts down the server")
	return s.http.Shutdown(ctx)
}
