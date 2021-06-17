package pgstore

import (
	"database/sql"
	"errors"

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
		err = res.Scan(&crypt.ID, &crypt.Title, &crypt.Ticker, &crypt.CreatedAt)
		if err != nil {
			return nil, err
		}
		*crypto_instruments = append(*crypto_instruments, crypt)
	}

	return crypto_instruments, nil
}

func (c *CryptoInstrumentRepository) GetPopularCryptoInstrumentByUserID(id int64) (*models.CryptoInstrument, error) {

	q := `SELECT * FROM crypto_instruments WHERE id=(
		SELECT o.crypto_instrument_id
		FROM crypto_deals o
		  LEFT JOIN crypto_deals b
			  ON o.crypto_instrument_id > b.crypto_instrument_id
		WHERE b.crypto_instrument_id is NULL AND o.user_id=$1
		LIMIT 1)`

	res, err := c.db.Query(q, id)
	if res == nil {
		return nil, errors.New("Deals not found")
	}
	if err != nil {
		return nil, err
	}

	instrument := &models.CryptoInstrument{}

	for res.Next() {

		err = res.Scan(&instrument.ID, &instrument.Title, &instrument.Ticker, &instrument.CreatedAt)
		if err != nil {
			return nil, err
		}

	}

	return instrument, nil
}

func (c *CryptoInstrumentRepository) GetPopularCryptoInstrumentsID() ([]int64, error) {

	q := "SELECT crypto_instrument_id, COUNT(id) AS i FROM crypto_deals GROUP BY crypto_instrument_id ORDER BY i desc LIMIT 5"

	rows, err := c.db.Query(
		q,
	)
	if err != nil {
		return nil, err
	}

	var ids []int64

	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func (c *CryptoInstrumentRepository) GetCryptoInstrumentByID(id int64) (*models.CryptoInstrument, error) {

	q := "SELECT * FROM crypto_instruments WHERE id=$1"

	res, err := c.db.Query(q, id)
	if err != nil {
		return nil, err
	}

	instrument := &models.CryptoInstrument{}

	for res.Next() {

		err = res.Scan(&instrument.ID, &instrument.Title, &instrument.Ticker, &instrument.CreatedAt)
		if err != nil {
			return nil, err
		}

	}

	return instrument, nil
}
