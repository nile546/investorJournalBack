package models

// Position ...
type Positions int8

const (
	// Long ...
	Long Positions = iota + 1

	// Short ...
	Short
)

// TimeFrame ...
type TimeFrames int8

const (
	// Min1 ...
	M1 TimeFrames = iota + 1

	// M2 ...
	M2

	// M3 ...
	M3

	// M5 ...
	M5

	// M10 ...
	M10

	// M15 ...
	M15

	// H1 ...
	H1

	// H2 ...
	H2

	// H4 ...
	H4

	// Day1 ...
	D1

	// Week1 ...
	W1

	// Month1 ...
	MN1
)

type Currencies int8

const (
	Usd Currencies = iota + 1

	Eur

	Rub
)
