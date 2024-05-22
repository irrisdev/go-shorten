package models

import "time"

type URL struct {
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	Clicks      int       `json:"clicks"`
	Creation    time.Time `json:"creation"`
	Expiration  time.Time `json:"expiration,omitempty"`
}
