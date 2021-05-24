package investInstruments

type Instruments interface {
	Stocks() StockInstrument
	Cryptos() CryptoInstrument
	Deposits() DepositInstruments
}

type StockInstrument interface {
	Grab() error
}

type CryptoInstrument interface {
}
type DepositInstruments interface {
}
