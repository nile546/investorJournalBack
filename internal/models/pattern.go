package models

import "time"

type Type int8

const (
	Stock Type = iota + 1
	Crypto
)

type Pattern struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Icon        *string   `json:"icon"`
	UserID      *int64    `json:"user_id"`
	Type        Type      `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
}
