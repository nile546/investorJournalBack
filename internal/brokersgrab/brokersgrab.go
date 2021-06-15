package brokersgrab

import "github.com/nile546/diplom/internal/models"

type Grab interface {
	TinkoffGrab() TinkoffGrab
}

type TinkoffGrab interface {
	GetTinkoffStockDeals(string) (*[]models.TinkoffOperation, error)
}
