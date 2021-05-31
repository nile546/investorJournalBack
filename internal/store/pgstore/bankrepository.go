package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type BankRepository struct {
	db *sql.DB
}

func (s *BankRepository) InsertBanks(*[]models.Bank) error {

	return nil
}
