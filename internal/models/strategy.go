package models

import "time"

type Strategy struct {
	ID             int64           `json:"id"`
	Name           *string         `json:"name"`
	Description    *string         `json:"description"`
	UserID         *int64          `json:"userId,omitempty"`
	InstrumentType InstrumentTypes `json:"instrumentType"`
	CreatedAt      time.Time       `json:"createdAt"`
}

func (s *Strategy) GetID() *int64 {
	if s == nil {
		return nil
	}

	return &s.ID
}
