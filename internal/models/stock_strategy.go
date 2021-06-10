package models

import "time"

type StockStrategy struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserID      int64     `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}
