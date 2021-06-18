package pgstore

import (
	"database/sql"
	"errors"
	"math"

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

	count := 0

	for res.Next() {
		count++
		bank := models.BankInstrument{}
		err = res.Scan(&bank.ID, &bank.Title, &bank.CreatedAt)
		if err != nil {
			return nil, err
		}
		*banks_instruments = append(*banks_instruments, bank)
	}

	if count == 0 {
		return nil, nil
	}

	return banks_instruments, nil
}

func (b *BankInstrumentRepository) GetPopularBankInstrumentByUserID(id int64) (*models.BankInstrument, error) {

	q := `SELECT * FROM banks_instruments WHERE id=(
		SELECT o.bank_instrument_id
		FROM deposit_deals o
		  LEFT JOIN deposit_deals b
			  ON o.bank_instrument_id > b.bank_instrument_id
		WHERE b.bank_instrument_id is NULL AND o.user_id=$1
		LIMIT 1)`

	instrument := &models.BankInstrument{}

	err := b.db.QueryRow(q, id).Scan(
		&instrument.ID,
		&instrument.Title,
		&instrument.CreatedAt)
	if err != nil {
		return nil, err
	}

	return instrument, nil
}

func (b *BankInstrumentRepository) GetPopularBankInstrumentsID() ([]int64, error) {

	q := "SELECT bank_instrument_id, COUNT(id) AS i FROM deposit_deals GROUP BY bank_instrument_id ORDER BY i desc LIMIT 5"

	rows, err := b.db.Query(
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

func (b *BankInstrumentRepository) GetBankInstrumentByID(id int64) (*models.BankInstrument, error) {

	q := "SELECT * FROM banks_instruments WHERE id=$1"

	instrument := &models.BankInstrument{}

	err := b.db.QueryRow(q, id).Scan(
		&instrument.ID,
		&instrument.Title,
		&instrument.CreatedAt)
	if err != nil {
		return nil, err
	}

	return instrument, nil
}

func (b *BankInstrumentRepository) GetAll(tp *models.TableParams) error {

	queryParams := []interface{}{
		tp.Pagination.ItemsPerPage,
		tp.Pagination.PageNumber * tp.Pagination.ItemsPerPage,
	}

	var like string
	if tp.SearchText != "" {
		queryParams = append(queryParams, tp.SearchText+"%")
		like = `
		WHERE title LIKE $3`
	}

	q := `SELECT id, title, created_at
	FROM banks_instruments ` + like +
		` LIMIT $1 
	OFFSET $2;`

	rows, err := b.db.Query(
		q,
		queryParams...,
	)
	if err != nil {
		return err
	}

	defer rows.Close()

	source := []models.BankInstrument{}

	count := 0

	for rows.Next() {
		count++
		var si models.BankInstrument
		err = rows.Scan(&si.ID, &si.Title, &si.CreatedAt)
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
	FROM banks_instruments
	`
	var itemsCount int
	if err = b.db.QueryRow(q).Scan(&itemsCount); err != nil {
		return err
	}
	defer rows.Close()

	tp.Pagination.PageCount = 0

	if itemsCount > 0 {
		tp.Pagination.PageCount = int(math.Ceil(float64(itemsCount) / float64(tp.Pagination.ItemsPerPage)))
	}

	return nil

}
