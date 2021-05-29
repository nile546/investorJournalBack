package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type StockRepository struct {
	db *sql.DB
}

func (s *StockRepository) InsertStocks(*[]models.Stock) error {

	return nil
}
