package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type CryptoInstrumentRepository struct {
	db *sql.DB
}

func (c *CryptoInstrumentRepository) InsertCryptoInstruments(cryptos *[]models.CryptoInstrument) (err error) {

	if len(*cryptos) == 0 {
		return
	}

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

func (c *CryptoInstrumentRepository) GetAllCryptoInstruments() (*[]models.CryptoInstrument, error) {

	q := `SELECT * FROM crypto_instruments`

	res, err := c.db.Query(q)
	if err != nil {
		return nil, err
	}

	crypto_instruments := &[]models.CryptoInstrument{}

	for res.Next() {
		crypt := models.CryptoInstrument{}
		err = res.Scan(&crypt.ID, &crypt.Title, &crypt.Ticker)
		if err != nil {
			return nil, err
		}
		*crypto_instruments = append(*crypto_instruments, crypt)
	}

	return crypto_instruments, nil
}
