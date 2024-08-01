package handlers

import (
	"encoding/json"
	"net/http"

	"jeorozco.com/go/url-shortener/models"
)

type URLResponse struct {
	ShortURL string `json:"short_url"`
}

func CreateURL(
	w http.ResponseWriter,
	r *http.Request,
) {
	var url models.LongURL
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if url.Url == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	shortURL := models.New(url)

	models.CacheMutex.Lock()
	models.UrlCache[shortURL.ID] = shortURL
	models.CacheMutex.Unlock()

	response := URLResponse{
		ShortURL: shortURL.NewURL,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
