package main

import (
	"context"
	flag "github.com/spf13/pflag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a context that listens for SIGINT (Ctrl+C) or SIGTERM (docker/kubernetes/process manager stop)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var (
		host = flag.String("host", "localhost", "server host")
		port = flag.Int("port", 25565, "server port")
	)

	slog.Info("parameters", slog.String("host", *host), slog.Int("port", *port))

	// Block until a signal is received to terminate the program
	<-ctx.Done()

	// Timeout for cleaning up
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Cleanup code

	slog.Info("cleanup complete, existing cleanly")

}
