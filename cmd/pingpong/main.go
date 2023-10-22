package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pingpong/internal/handler"
	"pingpong/internal/middleware"
	"time"
)

// main is the entry point of the "pingpong" application.
// It sets up an HTTP server, starts it, and gracefully handles shutdown
// when an interrupt signal is received.
func main() {
	mux := http.NewServeMux()
	pingServer := &handler.PingServer{}
	healthServer := &handler.HealthServer{}

	mux.Handle("/ping", pingServer)
	mux.Handle("/health", healthServer)

	// Wrap our server with the logging middleware
	loggingMux := middleware.Logging(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: loggingMux,
	}

	// Run the server in a goroutine to allow it to be non-blocking
	// and to listen for shutdown signals concurrently.
	go func() {
		log.Println("Server started on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server start failed: %v", err)
		}
	}()

	// Create a channel to listen for interrupt signals for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Block until an interrupt is received
	<-stop

	// Start graceful shutdown with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%v", err)
	}
}
