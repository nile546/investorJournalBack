package brokersgrab

import (
	"time"

	"github.com/nile546/diplom/internal/models"
)

type Grab interface {
	TinkoffGrab() TinkoffGrab
}

type TinkoffGrab interface {
	GetTinkoffStockDeals(string, time.Time) (*[]models.TinkoffOperation, error)
}
