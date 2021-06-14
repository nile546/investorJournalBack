package models

import "time"

//"Position", "TimeFrame", "Currency" declared in stock deals model

type CryptoDeal struct {
	ID              int64            `json:"id"`
	Crypto          CryptoInstrument `json:"crypto"`
	Currency        *Currency        `json:"currency"`
	Strategy        *Strategy        `json:"strategy"`
	Pattern         *Pattern         `json:"pattern"`
	Position        *Position        `json:"position"`
	TimeFrame       *TimeFrame       `json:"time_frame"`
	EnterDateTime   time.Time        `json:"enter_datetime"`
	EnterPoint      int64            `json:"enter_point"`
	StopLoss        *int64           `json:"stop_loss"`
	Quantity        int              `json:"quantity"`
	ExitDateTime    *time.Time       `json:"exit_datetime"`
	ExitPoint       *int64           `json:"exit_point"`
	RiskRatio       float32          `json:"risk_ratio"`
	Result          *int64           `json:"result"`
	ResultInPercent float64          `json:"result_in_percent"`
	StartDeposit    int64            `json:"start_deposit"`
	EndDeposit      int64            `json:"end_deposit"`
	UserID          int64            `json:"user_id"`
}
