package handlers

import (
	"encoding/json"
	"net/http"

	"jeorozco.com/go/url-shortener/models"
)

func GetURL(w http.ResponseWriter,
	r *http.Request) {
	urlID := r.PathValue("id")
	models.CacheMutex.RLock()
	url, ok := models.UrlCache[urlID]
	models.CacheMutex.RUnlock()
	if !ok {
		http.Error(w, "url not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(url.OriginalURL)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
