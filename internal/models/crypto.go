package models

import "time"

type Crypto struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Ticker    string    `json:"ticker"`
	CreatedAt time.Time `json:"createdAt"`
}
