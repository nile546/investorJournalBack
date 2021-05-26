package investInstruments

import "github.com/nile546/diplom/internal/models"

type Instruments interface {
	Stocks() StockInstrument
	Cryptos() CryptoInstrument
	Deposits() DepositInstruments
}

type StockInstrument interface {
	GrabPage() ([]*models.Stock, error)
}

type CryptoInstrument interface {
}
type DepositInstruments interface {
}
