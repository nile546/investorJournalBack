package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type StockDealPartRepository struct {
	db *sql.DB
}

func (s *StockDealPartRepository) InsertStockDealPart(part *models.StockDealParts) error {

	return nil
}
