// models/urlModel.go
package models

import (
	"time"

	"jeorozco.com/go/url-shortener/db"
	"jeorozco.com/go/url-shortener/services"

	"github.com/google/uuid"
)

type ShortURL struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	NewURL      string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type LongURL struct {
	Url string `json:"url"`
}

func New(url LongURL) (ShortURL, error) {
	encoded := services.UUIDToShortID(uuid.New())
	baseURL := "https://www.nanourl-dev.xyz" // Update this with your fly.io domain
	newURL := baseURL + "/" + encoded

	shortURL := ShortURL{
		ID:          encoded,
		OriginalURL: url.Url,
		NewURL:      newURL,
		CreatedAt:   time.Now().UTC(),
	}

	// Insert into database
	_, err := db.DB.Exec(
		`INSERT INTO urls (id, original_url, new_url, created_at) 
		 VALUES ($1, $2, $3, $4)`,
		shortURL.ID, shortURL.OriginalURL, shortURL.NewURL, shortURL.CreatedAt,
	)
	if err != nil {
		return ShortURL{}, err
	}

	return shortURL, nil
}

func GetByID(id string) (ShortURL, error) {
	var url ShortURL
	err := db.DB.QueryRow(
		`SELECT id, original_url, new_url, created_at 
		 FROM urls WHERE id = $1`,
		id,
	).Scan(&url.ID, &url.OriginalURL, &url.NewURL, &url.CreatedAt)
	return url, err
}
