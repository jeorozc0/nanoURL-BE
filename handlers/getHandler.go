package handlers

import (
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
	_, err := http.Get(url.OriginalURL)
	if err != nil {
		http.Error(w, "url does not exist", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusMovedPermanently)
}
