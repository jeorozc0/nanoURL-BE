package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/rs/cors"
	"jeorozco.com/go/url-shortener/db"
	"jeorozco.com/go/url-shortener/handlers"
	"jeorozco.com/go/url-shortener/middleware"
)

// getEnv retrieves environment variable or returns default value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// getCORSOrigins splits the CORS_ORIGINS environment variable into a slice
func getCORSOrigins() []string {
	origins := getEnv("CORS_ORIGINS", "https://www.nanourl-dev.xyz")
	return strings.Split(origins, ",")
}

func main() {
	// Initialize database
	db.InitDB()
	defer db.DB.Close()

	// Create new mux
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("POST /url", handlers.CreateURL)
	mux.HandleFunc("GET /{id}", handlers.GetURL)

	// Configure CORS
	corsOptions := cors.Options{
		AllowedOrigins: getCORSOrigins(),
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
		},
		AllowCredentials: true,
		// Enable Debugging for testing
		Debug: getEnv("CORS_DEBUG", "false") == "true",
	}

	// Create handler chain
	handler := cors.New(corsOptions).Handler(mux)
	loggingHandler := middleware.Logging(handler)

	// Configure server
	port := getEnv("PORT", "8080")
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      loggingHandler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server
	go func() {
		log.Printf("Server starting on port %s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server is shutting down...")

	// Create shutdown deadline
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error during server shutdown: %v\n", err)
	}

	log.Println("Server exited properly")
}
