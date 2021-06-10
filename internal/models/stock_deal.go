package models

import "time"

// Position ...
type Position int8

const (
	// Long ...
	Long Position = iota + 1

	// Short ...
	Short
)

// TimeFrame ...
type TimeFrame int8

const (
	// M1 ...
	M1 TimeFrame = iota + 1

	// M2 ...
	M2

	// M3 ...
	M3

	// M5 ...
	M5

	// M10 ...
	M10

	// M15 ...
	M15

	//TODO: Need to add all timeframe. Gets from tinkoff terminal
)

type StockDeal struct {
	ID              int64           `json:"id"`
	Stock           StockInstrument `json:"stock"`
	Strategy        StockStrategy   `json:"strategy"`
	Pattern         StockPattern    `json:"pattern"`
	Position        Position        `json:"position"`
	TimeFrame       TimeFrame       `json:"time_frame"`
	EnterDateTime   time.Time       `json:"enter_datetime"`
	EnterPoint      int64           `json:"enter_point"`
	StopLoss        int64           `json:"stop_loss"`
	Quantity        int             `json:"quantity"`
	ExitDateTime    *time.Time      `json:"exit_datetime"`
	ExitPoint       int64           `json:"exit_point"`
	RiskRatio       float32         `json:"risk_ratio"`
	Result          int64           `json:"result"`
	ResultInPercent float64         `json:"result_in_percent"`
	StartDeposit    int64           `json:"start_deposit"`
	EndDeposit      int64           `json:"end_deposit"`
}
