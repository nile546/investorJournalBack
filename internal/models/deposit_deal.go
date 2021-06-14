package models

import "time"

type DepositDeal struct {
	ID            int64          `json:"id"`
	Bank          BankInstrument `json:"bank"`
	Strategy      *Strategy      `json:"strategy"`
	EnterDateTime time.Time      `json:"enter_datetime"`
	Percent       float64        `json:"percent"`
	ExitDateTime  time.Time      `json:"exit_datetime"`
	StartDeposit  int64          `json:"start_deposit"`
	EndDeposit    int64          `json:"end_deposit"`
	UserID        int64          `json:"user_id"`
	//Pattern       *Pattern       `json:"pattern"`
	//Position      *Position      `json:"position"`
	//TimeFrame     *TimeFrame     `json:"time_frame"`
	//Deposit       int64          `json:"enter_point"`
	//ExitPoint     *int64         `json:"exit_point"`
	//RiskRatio       float32        `json:"risk_ratio"`
	//Result          *int64  `json:"result"`
	//ResultInPercent float64 `json:"result_in_percent"`

}
