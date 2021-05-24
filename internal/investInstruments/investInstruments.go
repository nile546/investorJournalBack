package investInstruments

type Instruments interface {
	Stocks() StockInstrument
	Cryptos() CryptoInstrument
	Banks() BankInstrument
}

type StockInstrument interface {
	grab() error
}

type CryptoInstrument interface {
}

type BankInstrument interface {
}
