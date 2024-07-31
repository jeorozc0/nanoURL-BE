package main

import (
	"fmt"
	"net/http"

	"jeorozco.com/go/url-shortener/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleRoot)

	fmt.Println("Server listening to :8080")
	http.ListenAndServe(":8080", mux)
}
