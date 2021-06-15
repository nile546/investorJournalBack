package models

import "time"

type StockInstrument struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Ticker    *string   `json:"ticker"`
	Type      *string   `json:"type"`
	Isin      *string   `json:"isin"`
	CreatedAt time.Time `json:"createdAt"`
}
