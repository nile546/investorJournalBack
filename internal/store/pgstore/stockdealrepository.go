package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type StockDealRepository struct {
	db *sql.DB
}

func (r *StockDealRepository) GetAll(tp *models.TableParams) error {

	q := `
	SELECT sd.id 
	FROM stock_deals AS sd
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

	source := []models.StockDeal{}

	for rows.Next() {
		var sd models.StockDeal
		err = rows.Scan(&sd.ID)
		if err != nil {
			return err
		}

		source = append(source, sd)
	}

	tp.Source = source
	return nil
}

func (r *StockDealRepository) GetStockDealsIDByISIN(ISIN string) (int64, error) {
	q := `SELECT id FROM stock_deals 
	WHERE (stock_instrument_id= 
	SELECT id FROM stocks_instruments 
	WHERE isin=$1) 
	AND (exit_datetime IS NULL) AND (variability=false)`

	var stock_dealID int64

	if err := r.db.QueryRow(q, ISIN).Scan(&stock_dealID); err != nil {
		return 0, err
	}

	return stock_dealID, nil
}
