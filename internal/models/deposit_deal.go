package models

import "time"

type DepositDeal struct {
	ID            int64          `json:"id"`
	Bank          BankInstrument `json:"bank"`
	Currency      *Currencies    `json:"currency"`
	EnterDateTime time.Time      `json:"enterDatetime"`
	Percent       float64        `json:"percent"`
	ExitDateTime  *time.Time     `json:"exitDatetime,omitempty"`
	StartDeposit  *int64         `json:"startDeposit,omitempty"`
	EndDeposit    *int64         `json:"endDeposit,omitempty"`
	Result        *int64         `json:"result,omitempty"`
}
