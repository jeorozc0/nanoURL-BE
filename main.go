package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/cors"
	"jeorozco.com/go/url-shortener/handlers"
	"jeorozco.com/go/url-shortener/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /url", handlers.CreateURL)
	mux.HandleFunc("GET /{id}", handlers.GetURL)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // Adjust this to your React app's URL
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(mux)
	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(handler),
	}

	go func() {
		fmt.Println("Server listening to :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server is shutting down...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout deadline
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v\n", err)
	}

	log.Println("Server exited properly")
}
