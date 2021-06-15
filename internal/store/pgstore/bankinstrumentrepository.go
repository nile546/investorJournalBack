package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type BankInstrumentRepository struct {
	db *sql.DB
}

func (b *BankInstrumentRepository) InsertBanksInstruments(banks *[]models.BankInstrument) (err error) {

	if len(*banks) == 0 {
		return
	}

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

func (b *BankInstrumentRepository) GetAllBankInstruments() (*[]models.BankInstrument, error) {

	q := `SELECT * FROM banks_instruments`

	res, err := b.db.Query(q)
	if err != nil {
		return nil, err
	}

	banks_instruments := &[]models.BankInstrument{}

	for res.Next() {
		bank := models.BankInstrument{}
		err = res.Scan(&bank.ID, &bank.Title)
		if err != nil {
			return nil, err
		}
		*banks_instruments = append(*banks_instruments, bank)
	}

	return banks_instruments, nil
}
