package models

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"jeorozco.com/go/url-shortener/services"
)

type ShortURL struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	NewURL      string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at"`
}

func New(url LongURL) ShortURL {
	encoded := services.UUIDToShortID(uuid.New())
	newURL := fmt.Sprintf("localhost:8080/%s", encoded)
	shortUrl := ShortURL{
		ID:          encoded,
		OriginalURL: url.Url,
		NewURL:      newURL,
		CreatedAt:   time.Now().UTC(),
	}
	return shortUrl
}

type LongURL struct {
	Url string `json:"url"`
}

var UrlCache = make(map[string]ShortURL)

var CacheMutex sync.RWMutex
