package models

import "time"

//"Position", "TimeFrame", "Currency" declared in stock deals model

type CryptoDeal struct {
	ID              int64            `json:"id"`
	Crypto          CryptoInstrument `json:"crypto"`
	Currency        *Currencies      `json:"currency"`
	Strategy        *Strategy        `json:"strategy"`
	Pattern         *Pattern         `json:"pattern"`
	Position        *Positions       `json:"position"`
	TimeFrame       *TimeFrames      `json:"timeFrame"`
	EnterDateTime   time.Time        `json:"enterDatetime"`
	EnterPoint      int64            `json:"enterPoint"`
	StopLoss        *int64           `json:"stopLoss"`
	Quantity        int              `json:"quantity"`
	ExitDateTime    *time.Time       `json:"exitDatetime"`
	ExitPoint       *int64           `json:"exitPoint"`
	RiskRatio       float32          `json:"riskRatio"`
	Result          *int64           `json:"result"`
	ResultInPercent float64          `json:"resultInPercent"`
	StartDeposit    int64            `json:"startDeposit"`
	EndDeposit      int64            `json:"endDeposit"`
}
