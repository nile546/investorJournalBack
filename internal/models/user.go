package models

import "time"

type user struct {
	Login     string    `json:"login"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	isActive  bool      `json:"isActive"`
	createdAt time.Time `json:"createdAt"`
}
