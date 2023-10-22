// Package httpserver implements HTTP server.
package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/vildan-valeev/perx_test/internal/config"
	"github.com/vildan-valeev/perx_test/internal/service"
	"github.com/vildan-valeev/perx_test/internal/transport/http/api"
)

type Server struct {
	http   *http.Server
	config config.Config
}

// New returns a new instance of Server.
func New(ctx context.Context, cfg config.Config, services *service.Services) *Server {
	s := &Server{
		config: cfg,
	}

	// version 1
	r := api.NewAPI(ctx, cfg, services.Item)

	s.http = &http.Server{
		Addr:    net.JoinHostPort(s.config.IP, s.config.HTTPPort),
		Handler: r,
	}

	return s
}

// Open validates the server options and begins listening on the bind address.
func (s *Server) Open() error {
	go func() {
		log.Printf("Start HTTP on %s:%s\n", s.config.IP, s.config.HTTPPort)

		if err := s.http.ListenAndServe(); err != nil {
			log.Println("failed to http serve (or closing)")
		}
	}()

	return nil
}

// Close gracefully shuts down the server.
func (s *Server) Close(ctx context.Context) error {
	log.Println("gracefully shuts down the server")
	return s.http.Shutdown(ctx)
}
