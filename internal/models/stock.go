package models

import "time"

type Stock struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Ticker    string    `json:"ticker"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
}
