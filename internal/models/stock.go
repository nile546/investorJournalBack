package models

import "time"

type Stock struct {
	ID        int64     `json:"id"`
	Ticket    string    `json:"ticket"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}
