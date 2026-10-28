package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	flag "github.com/spf13/pflag"
)

func run(ctx context.Context) error {
	var (
		host = flag.String("host", "localhost", "server host")
		port = flag.Uint64("port", 25565, "server port")
	)

	slog.Info("parameters", slog.String("host", *host), slog.Uint64("port", *port))

	listener, err := net.Listen("tcp", *host+":"+fmt.Sprintf("%d", *port))
	if err != nil {
		return fmt.Errorf("error listening: %w", err)
	}
	defer listener.Close()

	// Block until a signal is received to terminate the program
	<-ctx.Done()

	return nil
}

func main() {
	// Create a context that listens for SIGINT (Ctrl+C) or SIGTERM (docker/kubernetes/process manager stop)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx); err != nil {
		slog.Error("application terminated with error", slog.Any("error", err))
	}

	// Timeout for cleaning up
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Cleanup code

	slog.Info("cleanup complete, existing cleanly")

}
