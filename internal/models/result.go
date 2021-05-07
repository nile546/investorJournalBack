package models

type ResultStatuses int8

const (
	Ok ResultStatuses = iota + 1
	Error
)

type Result struct {
	Status       ResultStatuses `json:"status"`
	Payload      interface{}    `json:"payload,omitempty"`
	ErrorMessage string         `json:"errorMessage,omitempty"`
}
