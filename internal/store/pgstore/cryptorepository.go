package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type CryptoRepository struct {
	db *sql.DB
}

func (c *CryptoRepository) InsertCrypto(cryptos *[]models.Crypto) (err error) {

	q := `INSERT INTO crypto (title, ticker) VALUES `

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

func (c *CryptoRepository) TruncateCrypto() (err error) {

	q := `TRUNCATE TABLE crypto RESTART IDENTITY`

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
