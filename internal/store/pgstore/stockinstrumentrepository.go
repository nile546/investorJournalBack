package pgstore

import (
	"database/sql"
	"errors"
	"math"
	"strings"

	"github.com/nile546/diplom/internal/models"
)

type StockInstrumentRepository struct {
	db *sql.DB
}

func (s *StockInstrumentRepository) InsertStocksInstruments(stocks *[]models.StockInstrument) (err error) {

	if len(*stocks) == 0 {
		return
	}

	q := `INSERT INTO stocks_instruments (title, ticker, type, isin) VALUES `

	var res sql.Result
	var buf string

	for i, stock := range *stocks {
		if i == len(*stocks)-1 {
			buf = "('"
			buf += strings.Replace(stock.Title, "'", "''", -1) + "', '" + strings.Replace(*stock.Ticker, "'", "''", -1) + "', '" + strings.Replace(*stock.Type, "'", "''", -1) + "', '" + strings.Replace(*stock.Isin, "'", "''", -1)
			buf += "')"
			q += buf
			break
		}
		buf = "('"
		buf += strings.Replace(stock.Title, "'", "''", -1) + "', '" + strings.Replace(*stock.Ticker, "'", "''", -1) + "', '" + strings.Replace(*stock.Type, "'", "''", -1) + "', '" + strings.Replace(*stock.Isin, "'", "''", -1)
		buf += "'), "
		q += buf
	}

	res, err = s.db.Exec(q)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *StockInstrumentRepository) GetAllStockInstruments() (*[]models.StockInstrument, error) {

	q := `SELECT * FROM stocks_instruments`

	res, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}

	stocks_instruments := &[]models.StockInstrument{}

	for res.Next() {
		si := models.StockInstrument{}
		err = res.Scan(&si.ID, &si.Title, &si.Ticker, &si.Type, &si.Isin, &si.CreatedAt)
		if err != nil {
			return nil, err
		}
		*stocks_instruments = append(*stocks_instruments, si)
	}

	return stocks_instruments, nil
}

func (s *StockInstrumentRepository) GetInstrumentByISIN(ISIN string) (*models.StockInstrument, error) {
	q := `SELECT * FROM stocks_instruments WHERE isin=$1`

	res, err := s.db.Query(q, ISIN)
	if err != nil {
		return nil, err
	}

	instrument := &models.StockInstrument{
		Isin: &ISIN,
	}

	for res.Next() {

		err = res.Scan(&instrument.ID, &instrument.Title, &instrument.Ticker, &instrument.Type, &instrument.Isin, &instrument.CreatedAt)
		if err != nil {
			return nil, err
		}

	}

	return instrument, nil
}

func (s *StockInstrumentRepository) GetPopularStockInstrumentByUserID(id int64) (*models.StockInstrument, error) {

	q := `SELECT * FROM stocks_instruments WHERE id=(
		SELECT o.stock_instrument_id
		FROM stock_deals o
		  LEFT JOIN stock_deals b
			  ON o.stock_instrument_id > b.stock_instrument_id
		WHERE b.stock_instrument_id is NULL AND o.user_id=$1
		LIMIT 1)`

	instrument := &models.StockInstrument{}

	err := s.db.QueryRow(q, id).Scan(
		&instrument.ID,
		&instrument.Title,
		&instrument.Ticker,
		&instrument.Type,
		&instrument.Isin,
		&instrument.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return instrument, nil
}

func (s *StockInstrumentRepository) GetPopularStockInstrumentsID() ([]int64, error) {

	q := "SELECT stock_instrument_id, COUNT(id) AS i FROM stock_deals GROUP BY stock_instrument_id ORDER BY i desc LIMIT 5"

	rows, err := s.db.Query(
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

func (s *StockInstrumentRepository) GetStockInstrumentByID(id int64) (*models.StockInstrument, error) {

	q := "SELECT * FROM stocks_instruments WHERE id=$1"

	instrument := &models.StockInstrument{}

	err := s.db.QueryRow(q, id).Scan(
		&instrument.ID,
		&instrument.Title,
		&instrument.Ticker,
		&instrument.Type,
		&instrument.Isin,
		&instrument.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return instrument, nil
}

func (r *StockInstrumentRepository) GetAll(tp *models.TableParams) error {

	q := `SELECT id, title, ticker, type, isin, created_at
	FROM stocks_instruments
	LIMIT $1 
	OFFSET $2;
	`

	rows, err := r.db.Query(
		q,
		tp.Pagination.ItemsPerPage,
		tp.Pagination.PageNumber*tp.Pagination.ItemsPerPage,
	)
	if err != nil {
		return err
	}

	defer rows.Close()

	source := []models.StockInstrument{}

	count := 0

	for rows.Next() {
		count++
		var si models.StockInstrument
		err = rows.Scan(&si.ID, &si.Title, &si.Ticker, &si.Type, &si.Isin, &si.CreatedAt)
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
	FROM stocks_instruments
	`
	var itemsCount int
	if err = r.db.QueryRow(q).Scan(&itemsCount); err != nil {
		return err
	}
	defer rows.Close()

	tp.Pagination.PageCount = 0

	if count > 0 {
		tp.Pagination.PageCount = int(math.Ceil(float64(count) / float64(tp.Pagination.ItemsPerPage)))
	}

	return nil
}
