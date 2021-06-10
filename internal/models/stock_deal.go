package models

import "time"

type StockDeal struct {
	ID              int64      `json:"id"`
	Stock           int64      `json:"title_instrument"`
	Strategy        int64      `json:"strategy"`
	Pattern         int64      `json:"pattern"`
	Position        int64      `json:"position"`
	TimeFrame       int64      `json:"time_frame"`
	EnterDateTime   time.Time  `json:"enter_datetime"`
	EnterPoint      int64      `json:"enter_point"`
	StopLoss        int64      `json:"stop_loss"`
	Count           int64      `json:"count"`
	ExitDateTime    *time.Time `json:"exit_datetime"`
	ExitPoint       int64      `json:"exit_point"`
	RiskRatio       float64    `json:"risk_ratio"`
	Result          int64      `json:"result"`
	ResultInPercent float64    `json:"result_in_percent"`
}
