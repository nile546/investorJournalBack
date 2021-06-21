package models

import "time"

type Pattern struct {
	ID             int64           `json:"id"`
	Name           *string         `json:"name"`
	Description    *string         `json:"description"`
	Icon           *string         `json:"icon"`
	UserID         *int64          `json:"userId"`
	InstrumentType InstrumentTypes `json:"instrumentType"`
	CreatedAt      time.Time       `json:"createdAt"`
}

func (p *Pattern) GetID() *int64 {
	if p == nil {
		return nil
	}

	return &p.ID
}
