package investinstruments

import "github.com/nile546/diplom/internal/models"

type Instruments interface {
	Stocks() Stockinstrument
	Cryptos() Cryptoinstrument
	Deposits() Depositinstruments
}

type Stockinstrument interface {
	SPBGrab(u string) (*[]models.Stock, error)
	MSKGrab(u string) (*[]models.Stock, error)
}

type Cryptoinstrument interface {
}
type Depositinstruments interface {
}
