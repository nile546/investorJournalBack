package models

import "time"

type TinkoffOperation struct {
	ID        int64     `json:"id"`
	FIGI      string    `json:"figi"`
	Currency  int8      `json:"currency"`
	Quantity  int       `json:"quantity"`
	DateTime  time.Time `json:"enter_datetime"`
	Price     int64     `json:"enter_point"`
	Operation string    `json:"operation"`
}
