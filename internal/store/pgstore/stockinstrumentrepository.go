package pgstore

import (
	"database/sql"
	"strings"

	"github.com/nile546/diplom/internal/models"
)

type StockInstrumentRepository struct {
	db *sql.DB
}

func (s *StockInstrumentRepository) InsertStocksInstruments(stocks *[]models.StockInstrument) (err error) {

	q := `INSERT INTO stocks_instruments (title, ticker, type) VALUES `

	var res sql.Result
	var buf string

	for i, stock := range *stocks {
		if i == len(*stocks)-1 {
			buf = "('"
			buf += strings.Replace(stock.Title, "'", "''", -1) + "', '" + strings.Replace(*stock.Ticker, "'", "''", -1) + "', '" + strings.Replace(*stock.Type, "'", "''", -1)
			buf += "')"
			q += buf
			break
		}
		buf = "('"
		buf += strings.Replace(stock.Title, "'", "''", -1) + "', '" + strings.Replace(*stock.Ticker, "'", "''", -1) + "', '" + strings.Replace(*stock.Type, "'", "''", -1)
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

func (s *StockInstrumentRepository) TruncateStocksInstruments() (err error) {

	q := `TRUNCATE TABLE stocks_instruments RESTART IDENTITY`

	var res sql.Result

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
