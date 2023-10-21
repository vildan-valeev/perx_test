package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vildan-valeev/perx_test/internal/config"
	"github.com/vildan-valeev/perx_test/internal/repository"
	"github.com/vildan-valeev/perx_test/internal/service"
	"github.com/vildan-valeev/perx_test/internal/transport/server"
)

func main() {
	workersCount := flag.Uint("n", 4, "workers count")
	flag.Parse()

	log.Printf("Start App with %d workers\n", *workersCount)
	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

	go func() {
		sig := <-sigs

		log.Printf("Shutting down server. Reason: %s...\n", sig.String())

		cancel()
	}()

	// Instantiate a new type to represent our application.
	m := NewMain(workersCount)

	// Execute program.
	if err := m.Run(ctx); err != nil {
		log.Println("Run server error")

		_ = m.Close(ctx)

		os.Exit(1)
	}

	// Wait for CTRL-C.
	<-ctx.Done()

	// Clean up program.
	if err := m.Close(ctx); err != nil {
		log.Println("Shutting down server error")
		os.Exit(1)
	}

	log.Println("Bye!")
}

// Main represents the program.
type Main struct {
	// Config parsed config data.
	Config *config.Config
	// HTTP server for handling communication.
	Srv *server.Server
}

// NewMain returns a new instance of Main.
func NewMain(n *uint) *Main {
	log.Println("Init config")

	cfg := config.New(n)

	return &Main{
		Config: cfg,
	}
}

// Run executes the program. The configuration should already be set up before
// calling this function.
func (m *Main) Run(ctx context.Context) (err error) {
	repositories := repository.NewRepositories()
	services := service.NewServices(service.Deps{
		Repos: repositories,
	})

	m.Srv = server.New(ctx, *m.Config, services)

	// Start the server.
	return m.Srv.Open()
}

// Close gracefully stops the program.
func (m *Main) Close(ctx context.Context) (err error) { //nolint
	if m.Srv != nil {
		_ = m.Srv.Close(ctx)
	}

	return nil
}
