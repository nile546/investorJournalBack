package models

import "time"

type TinkoffOperation struct {
	ID        int64      `json:"id"`
	ISIN      string     `json:"isin"`
	Currency  Currencies `json:"currency"`
	Quantity  int        `json:"quantity"`
	DateTime  time.Time  `json:"enterDatetime"`
	Price     int64      `json:"enterPoint"`
	Operation Type       `json:"operation"`
}
