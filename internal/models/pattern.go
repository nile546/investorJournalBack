package models

import "time"

type TypePattern int8

const (
	Stock TypePattern = iota + 1
	Crypto
)

type Pattern struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Description *string     `json:"description"`
	Icon        *string     `json:"icon"`
	UserID      *int64      `json:"user_id"`
	Type        TypePattern `json:"type"`
	CreatedAt   time.Time   `json:"created_at"`
}
