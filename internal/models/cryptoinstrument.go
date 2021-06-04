package models

import "time"

type CryptoInstrument struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Ticker    string    `json:"ticker"`
	CreatedAt time.Time `json:"createdAt"`
}
