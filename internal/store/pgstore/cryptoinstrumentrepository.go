package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type CryptoInstrumentsRepository struct {
	db *sql.DB
}

func (c *CryptoInstrumentsRepository) InsertCryptoInstruments(cryptos *[]models.CryptoInstrument) (err error) {

	q := `INSERT INTO crypto_instruments (title, ticker) VALUES `

	var res sql.Result
	var buf string

	for i, crypto := range *cryptos {
		if i == len(*cryptos)-1 {
			buf = "('"
			buf += crypto.Title + "', '" + crypto.Ticker
			buf += "')"
			q += buf
			break
		}
		buf = "('"
		buf += crypto.Title + "', '" + crypto.Ticker
		buf += "'), "
		q += buf
	}

	res, err = c.db.Exec(q)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (c *CryptoInstrumentsRepository) TruncateCryptoInstruments() (err error) {

	q := `TRUNCATE TABLE crypto_instruments RESTART IDENTITY`

	var res sql.Result

	res, err = c.db.Exec(q)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
