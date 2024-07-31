package models

type ShortURL struct {
	ID          int64  `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	CreatedAt   string `json:"created_at"`
}

type LongURL struct {
	Url string `json:"url"`
}
