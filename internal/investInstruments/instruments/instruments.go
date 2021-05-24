package instruments

import (
	"github.com/nile546/diplom/internal/investInstruments"
)

type Instruments struct {
	stockInstrumnet  *StockInstrument
	bankInstrumnet   *BankInstrument
	cryptoInstrumnet *CryptoInstrument
}

func New() *Instruments {
	return &Instruments{}
}

func (i *Instruments) Stocks() investInstruments.StockInstrument {

	if i.stockInstrumnet != nil {
		return i.stockInstrumnet
	}

	i.stockInstrumnet = &StockInstrument{}

	return i.stockInstrumnet
}

func (i *Instruments) Cryptos() investInstruments.CryptoInstrument {

	if i.cryptoInstrumnet != nil {
		return i.cryptoInstrumnet
	}

	i.cryptoInstrumnet = &CryptoInstrument{}

	return i.cryptoInstrumnet
}

func (i *Instruments) Banks() investInstruments.BankInstrument {

	if i.bankInstrumnet != nil {
		return i.bankInstrumnet
	}

	i.bankInstrumnet = &BankInstrument{}

	return i.bankInstrumnet
}
