package models 

type Type int8

const(
	
	Sell Type = iota + 1

	Buy
)

type StockDealParts struct{
	StockDealID int64 `json:"stockdeal_id"`
	Quantity int64 `json:"quantity"`
	Type Type `json:"type"`
	Price int64 `json:"price"`
	DateTime time.Time `json:"date_time"`
}