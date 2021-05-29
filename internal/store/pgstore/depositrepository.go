package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type DepositRepository struct {
	db *sql.DB
}

func (s *DepositRepository) InsertBanks(*[]models.Bank) error {

	return nil
}
