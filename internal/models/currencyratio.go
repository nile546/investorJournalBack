package models

type CurrencyRatio struct {
	FirstCurrency  Currencies `json:"first_currency"`
	SecondCurrency Currencies `json:"second_currency"`
	Ratio          string     `json:"ratio"`
}
