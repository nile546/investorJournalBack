package models

import "time"

type Type int8

const (
	Buy Type = iota + 1

	Sell
)

type StockDealParts struct {
	ID          int64     `json:"id"`
	Quantity    int       `json:"quantity"`
	Type        Type      `json:"type"`
	Price       int64     `json:"price"`
	DateTime    time.Time `json:"dateTime"`
	StockDealId int64     `json:"stockdealId"`
}
