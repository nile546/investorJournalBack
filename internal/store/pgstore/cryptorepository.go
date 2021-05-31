package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type CryptoRepository struct {
	db *sql.DB
}

func (s *CryptoRepository) InsertCrypto(*[]models.Crypto) error {

	return nil
}
