package investinstruments

import (
	"github.com/nile546/diplom/internal/models"
)

type Instruments interface {
	Stocks() Stockinstrument
	Cryptos() Cryptoinstrument
	Deposits() Depositinstruments
}

type Stockinstrument interface {
	GrabAll(spburl string, mskurl string) (*[]models.Stock, error)
}

type Cryptoinstrument interface {
	GrabCrypto() error
}
type Depositinstruments interface {
	GrabBanks(banksUrl string) (*[]models.Bank, error)
}
