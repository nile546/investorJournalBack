package models

import "time"

type DepositDeal struct {
	ID            int64          `json:"id"`
	Bank          BankInstrument `json:"bank"`
	Currency      *Currency      `json:"currency"`
	Strategy      *Strategy      `json:"strategy"`
	EnterDateTime time.Time      `json:"enter_datetime"`
	Percent       float64        `json:"percent"`
	ExitDateTime  time.Time      `json:"exit_datetime"`
	StartDeposit  int64          `json:"start_deposit"`
	EndDeposit    int64          `json:"end_deposit"`
	Result        *int64         `json:"result"`
	UserID        int64          `json:"user_id"`
}
