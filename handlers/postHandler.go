package handlers

import (
	"encoding/json"
	"net/http"

	"jeorozco.com/go/url-shortener/models"
)

func CreateURL(w http.ResponseWriter, r *http.Request) {
	var url models.LongURL
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		sendErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if url.Url == "" {
		sendErrorResponse(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortURL, err := models.New(url)
	if err != nil {
		sendErrorResponse(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}

	response := URLResponse{
		ShortURL: shortURL.NewURL,
	}

	sendJSONResponse(w, response, http.StatusCreated)
}
