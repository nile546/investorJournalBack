package pgstore

import (
	"database/sql"
	"errors"
	"math"

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

	count := 0

	for res.Next() {
		count++
		crypt := models.CryptoInstrument{}
		err = res.Scan(&crypt.ID, &crypt.Title, &crypt.Ticker, &crypt.CreatedAt)
		if err != nil {
			return nil, err
		}
		*crypto_instruments = append(*crypto_instruments, crypt)
	}

	if count == 0 {
		return nil, nil
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

	instrument := &models.CryptoInstrument{}

	err := c.db.QueryRow(q, id).Scan(
		&instrument.ID,
		&instrument.Title,
		&instrument.Ticker,
		&instrument.CreatedAt,
	)
	if err != nil {
		return nil, err
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

	if !rows.Next() {
		return nil, errors.New("Deals not found")
	}

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

	instrument := &models.CryptoInstrument{}

	err := c.db.QueryRow(q, id).Scan(
		&instrument.ID,
		&instrument.Title,
		&instrument.Ticker,
		&instrument.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return instrument, nil
}

func (c *CryptoInstrumentRepository) GetAll(tp *models.TableParams) error {

	q := `SELECT id, title, ticker, created_at
	FROM crypto_instruments
	LIMIT $1 
	OFFSET $2;
	`

	rows, err := c.db.Query(
		q,
		tp.Pagination.ItemsPerPage,
		tp.Pagination.PageNumber*tp.Pagination.ItemsPerPage,
	)
	if err != nil {
		return err
	}

	source := []models.CryptoInstrument{}

	count := 0

	for rows.Next() {
		count++
		var si models.CryptoInstrument
		err = rows.Scan(&si.ID, &si.Title, &si.Ticker, &si.CreatedAt)
		if err != nil {
			return err
		}

		source = append(source, si)
	}

	if count == 0 {
		return nil
	}

	tp.Source = source

	q = `SELECT COUNT(id)
	FROM crypto_instruments
	`
	var itemsCount int
	if err = c.db.QueryRow(q).Scan(&itemsCount); err != nil {
		return err
	}
	defer rows.Close()

	tp.Pagination.PageCount = 0

	if itemsCount > 0 {
		tp.Pagination.PageCount = int(math.Ceil(float64(itemsCount) / float64(tp.Pagination.ItemsPerPage)))
	}

	return nil
}
