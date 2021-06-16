package models

import "time"

//"Type" declared in pattern model

type Strategy struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	UserID      *int64    `json:"userId"`
	Type        Type      `json:"type"`
	CreatedAt   time.Time `json:"createdAt"`
}
