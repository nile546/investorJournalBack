package models

import "time"

type Stock struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Ticket    string    `json:"ticket"`
	CreatedAt time.Time `json:"createdAt"`
}
