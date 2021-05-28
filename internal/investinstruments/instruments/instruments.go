package instruments

import (
	"github.com/nile546/diplom/internal/investinstruments"
)

type Instruments struct {
	stockinstrumnet    *Stockinstrument
	depositinstruments *Depositinstruments
	cryptoinstrumnet   *Cryptoinstrument
}

func New() *Instruments {
	return &Instruments{}
}

func (i *Instruments) Stocks() investinstruments.Stockinstrument {

	if i.stockinstrumnet != nil {
		return i.stockinstrumnet
	}

	i.stockinstrumnet = &Stockinstrument{}

	return i.stockinstrumnet
}

func (i *Instruments) Cryptos() investinstruments.Cryptoinstrument {

	if i.cryptoinstrumnet != nil {
		return i.cryptoinstrumnet
	}

	i.cryptoinstrumnet = &Cryptoinstrument{}

	return i.cryptoinstrumnet
}

func (i *Instruments) Deposits() investinstruments.Depositinstruments {

	if i.depositinstruments != nil {
		return i.depositinstruments
	}

	i.depositinstruments = &Depositinstruments{}

	return i.depositinstruments
}
