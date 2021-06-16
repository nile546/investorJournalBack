package models

import "time"

type DepositDeal struct {
	ID            int64          `json:"id"`
	Bank          BankInstrument `json:"bank"`
	Currency      *Currencies    `json:"currency"`
	Strategy      *Strategy      `json:"strategy"`
	EnterDateTime time.Time      `json:"enterDatetime"`
	Percent       float64        `json:"percent"`
	ExitDateTime  time.Time      `json:"exitDatetime"`
	StartDeposit  int64          `json:"startDeposit"`
	EndDeposit    int64          `json:"endDeposit"`
	Result        *int64         `json:"result"`
	UserID        int64          `json:"userId"`
}
