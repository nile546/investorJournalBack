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

func (r *StockDealRepository) GetStockDealsIDByISIN(ISIN string) (int64, error) {
	q := `SELECT id FROM stock_deals WHERE (stock_instrument_id= 
	SELECT id FROM stocks_instruments WHERE isin=$1) AND (exit_datetime IS NULL) AND (variability=false)`

	var stock_dealID int64

	if err := r.db.QueryRow(q, ISIN).Scan(&stock_dealID); err != nil {
		return 0, err
	}

	return stock_dealID, nil
}

func (r *StockDealRepository) CreateStockDeal(deal *models.StockDeal) (int64, error) {

	q := `INSERT INTO stock_deal (stock_insrument_id, currency, strategy_id, pattern_id, 
	position, time_frame, enter_datetime, enter_point, stop_loss, quantity, 
	exit_datetime, exit_point, risk_ratio, variabiliry, user_id) 
	VALUES ($1, $2, $3) 
	RETURNING id`

	var idStockDeal int64

	err := r.db.QueryRow(q, deal.Stock.ID, deal.Currency, deal.Strategy.ID, deal.Pattern.ID, deal.Position, deal.TimeFrame, deal.EnterDateTime, deal.EnterPoint, deal.StopLoss, deal.Quantity, deal.ExitDateTime, deal.RiskRatio, deal.Variability, deal.UserID).Scan(&idStockDeal)
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
