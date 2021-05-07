package models

import "time"

type User struct {
	Login     string    `json:"login"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
}
