package models

type URL struct {
	OriginalURL string
	ShortURL    string
	Clicks      int
	Creation    int64
	Expiration  int64
}
