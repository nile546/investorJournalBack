package investinstruments

import (
	"github.com/nile546/diplom/internal/models"
)

type Instruments interface {
	Stocks() Stockinstrument
	Cryptos() Cryptoinstrument
	Banks() Bankinstruments
}

type Stockinstrument interface {
	GrabAll(spburl string, mskurl string) (*[]models.Stock, error)
}

type Cryptoinstrument interface {
	GrabAll(cryptoUrl string, cryptoKey string) (*[]models.Crypto, error)
}
type Bankinstruments interface {
	GrabAll(bankiUrl string) (*[]models.Bank, error)
}
