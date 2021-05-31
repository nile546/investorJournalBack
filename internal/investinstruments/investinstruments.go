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
	GrabCrypto(cryptoUrl string) (*[]models.Crypto, error)
}
type Bankinstruments interface {
	GrabBanks(banksUrl string) (*[]models.Bank, error)
}
