package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type BankInstrumentRepository struct {
	db *sql.DB
}

func (b *BankInstrumentRepository) InsertBanksInstruments(banks *[]models.BankInstrument) (err error) {

	q := `INSERT INTO banks_instruments (title) VALUES `

	var res sql.Result
	var buf string

	for i, bank := range *banks {
		if i == len(*banks)-1 {
			buf = "('"
			buf += bank.Title
			buf += "')"
			q += buf
			break
		}
		buf = "('"
		buf += bank.Title
		buf += "'), "
		q += buf
	}

	res, err = b.db.Exec(q)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (b *BankInstrumentRepository) TruncateBanksInstruments() (err error) {

	q := `TRUNCATE TABLE banks_instruments RESTART IDENTITY`

	var res sql.Result

	res, err = b.db.Exec(q)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
