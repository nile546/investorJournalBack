package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type BankRepository struct {
	db *sql.DB
}

func (s *BankRepository) InsertBanks(banks *[]models.Bank) (err error) {

	q := `INSERT INTO banks (title) VALUES ($1)`

	var res sql.Result

	for _, bank := range *banks {
		if res, err = s.db.Exec(q, bank.Title); err != nil {
			//TODO: ADD TO LOGER
			continue
		}
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
