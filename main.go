package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
	"jeorozco.com/go/url-shortener/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleRoot)

	mux.HandleFunc("POST /url", handlers.CreateURL)
	mux.HandleFunc("GET /{id}", handlers.GetURL)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // Adjust this to your React app's URL
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(mux)
	fmt.Println("Server listening to :8080")
	http.ListenAndServe(":8080", handler)
}
