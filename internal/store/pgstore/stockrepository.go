package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type StockRepository struct {
	db *sql.DB
}

func (s *StockRepository) InsertStocks(stocks *[]models.Stock) (err error) {

	q := `INSERT INTO stocks (title, ticker, type) VALUES ($1, $2, $3)`

	var res sql.Result

	for _, stock := range *stocks {
		if res, err = s.db.Exec(q, stock.Title, stock.Ticker, stock.Type); err != nil {
			//TODO: ADD TO LOGER
			continue
		}
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
