package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"jeorozco.com/go/url-shortener/models"
)

var urlCache = make(map[int]models.LongURL)

var cacheMutex sync.RWMutex

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
		http.Error(w, "url is requiered", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	urlCache[len(urlCache)+1] = url
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
