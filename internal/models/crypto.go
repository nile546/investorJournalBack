package models

import "time"

type Crypto struct {
	id        int64     `json:"id"`
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	createdAt time.Time `json:"createdAt"`
}
