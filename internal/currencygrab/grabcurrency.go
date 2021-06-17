package currency

import "github.com/nile546/diplom/internal/models"

type GrabsCurrency interface {
	GrabCbr() GrabCBR
}

type GrabCBR interface {
	GrabUsdEur() (*[]models.CurrencyRatio, error)
}
