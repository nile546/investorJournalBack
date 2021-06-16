package pgstore

import (
	"database/sql"
	"time"

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

func (r *StockDealRepository) GetStockDealsIDByISIN(ISIN string) int64 {
	q := `SELECT id FROM stock_deals WHERE (stock_instrument_id=
		(SELECT id FROM stocks_instruments WHERE isin=$1) AND (exit_datetime IS NULL) AND (variability=false))`

	var stock_dealID int64

	if err := r.db.QueryRow(q, ISIN).Scan(&stock_dealID); err != nil {
		return 0
	}

	return stock_dealID
}

func (r *StockDealRepository) CreateOpenStockDeal(deal *models.StockDeal) (int64, error) {

	q := `INSERT INTO stock_deals (stock_instrument_id, currency, enter_datetime, enter_point, quantity, variability, user_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) 
	RETURNING id`

	var idStockDeal int64

	err := r.db.QueryRow(q, deal.Stock.ID, &deal.Currency, deal.EnterDateTime, deal.EnterPoint, deal.Quantity, deal.Variability, deal.UserID).Scan(&idStockDeal)
	if err != nil {
		return 0, err
	}

	return idStockDeal, nil
}

func (r *StockDealRepository) UpdateQuantityStockDeal(idStockDeal int64, addQuontity int) error {

	q := "UPDATE stock_deals SET quantity=(quantity + $1) WHERE id=$2"

	res, err := r.db.Exec(q, addQuontity, idStockDeal)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *StockDealRepository) SetStockDealCompleted(exitDateTime time.Time, exitPoint int64, idStockDeal int64) error {

	q := "UPDATE stock_deals SET (exit_datetime, exit_point)=($1, $2) WHERE id=$3"

	res, err := r.db.Exec(q, exitDateTime, exitPoint, idStockDeal)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
