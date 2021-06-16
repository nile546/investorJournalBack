package models

import "time"

type DealTypes int8

const (
	Buy DealTypes = iota + 1

	Sell
)

type StockDealParts struct {
	ID          int64     `json:"id"`
	Quantity    int       `json:"quantity"`
	DealType    DealTypes `json:"dealType"`
	Price       int64     `json:"price"`
	DateTime    time.Time `json:"dateTime"`
	StockDealId int64     `json:"stockdealId"`
}
