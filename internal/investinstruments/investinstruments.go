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
	GrabAll(spburl string, mskurl string) (*[]models.StockInstrument, error)
}

type Cryptoinstrument interface {
	GrabAll(cryptoUrl string, cryptoKey string) (*[]models.CryptoInstrument, error)
}
type Bankinstruments interface {
	GrabAll(bankiUrl string) (*[]models.BankInstrument, error)
}
