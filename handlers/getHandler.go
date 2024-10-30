package handlers

import (
	"database/sql"
	"net/http"

	"jeorozco.com/go/url-shortener/models"
)

func GetURL(w http.ResponseWriter, r *http.Request) {
	urlID := r.PathValue("id")
	url, err := models.GetByID(urlID)
	if err != nil {
		if err == sql.ErrNoRows {
			sendErrorResponse(w, "URL not found", http.StatusNotFound)
			return
		}
		sendErrorResponse(w, "Server error", http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, url.OriginalURL, http.StatusOK)
}
