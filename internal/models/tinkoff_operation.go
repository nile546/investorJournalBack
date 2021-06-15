package models

import "time"

type TinkoffOperation struct {
	ID        int64     `json:"id"`
	ISIN      string    `json:"isin"`
	Currency  Currency  `json:"currency"`
	Quantity  int       `json:"quantity"`
	DateTime  time.Time `json:"enter_datetime"`
	Price     int64     `json:"enter_point"`
	Operation Type      `json:"operation"`
}
