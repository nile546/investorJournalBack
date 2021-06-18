package pgstore

import (
	"database/sql"
	"time"

	"github.com/nile546/diplom/internal/models"
)

type StockDealRepository struct {
	db *sql.DB
}

func (r *StockDealRepository) CreateStockDeal(deal *models.StockDeal) error {

	q := `INSERT INTO stock_deals
	(stock_instrument_id, currency, strategy_id, pattern_id, 
	position, time_frame, enter_datetime, enter_point, stop_loss, 
	quantity, exit_datetime, exit_point, risk_ratio, user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`

	res, err := r.db.Exec(q, deal.Stock.ID, deal.Currency, deal.Strategy.ID,
		deal.Pattern.ID, deal.Position, deal.TimeFrame, deal.EnterDateTime,
		deal.EnterPoint, deal.StopLoss, deal.Quantity, deal.ExitDateTime,
		deal.ExitPoint, deal.RiskRatio, deal.UserID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *StockDealRepository) UpdateStockDeal(deal *models.StockDeal) error {

	q := `UPDATE stock_deals SET (stock_instrument_id, currency, strategy_id, pattern_id, 
		position, time_frame, enter_datetime, enter_point, stop_loss, 
		quantity, exit_datetime, exit_point, risk_ratio)=($1, $2, $3, $4, $5, $6,
		$7, $8, $9, $10, $11, $12, $13)	WHERE id=$14`

	res, err := r.db.Exec(q, deal.Stock.ID, deal.Currency, deal.Strategy.ID,
		deal.Pattern.ID, deal.Position, deal.TimeFrame, deal.EnterDateTime,
		deal.EnterPoint, deal.StopLoss, deal.Quantity, deal.ExitDateTime,
		deal.ExitPoint, deal.RiskRatio, deal.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *StockDealRepository) DeleteStockDeal(id int64) error {

	q := "DELETE FROM stock_deals WHERE id=$1"

	res, err := r.db.Exec(q, id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *StockDealRepository) GetStockDealByID(id int64) (*models.StockDeal, error) {

	q := `SELECT stock_instrument_id, strategy_id, pattern_id, currency, 
	position, time_frame, enter_datetime, enter_point, stop_loss, 
	quantity, exit_datetime, exit_point, risk_ratio, variability, 
	user_id FROM stock_deals where id=$1`

	res, err := r.db.Query(q, id)
	if err != nil {
		return nil, err
	}

	deal := &models.StockDeal{
		ID: id,
	}

	for res.Next() {

		err = res.Scan(&deal.Stock.ID, &deal.Strategy.ID, &deal.Pattern.ID,
			&deal.Currency, &deal.Position, &deal.TimeFrame, &deal.EnterDateTime,
			&deal.EnterPoint, &deal.StopLoss, &deal.Quantity, &deal.ExitDateTime,
			&deal.ExitPoint, &deal.RiskRatio, &deal.Variability, &deal.UserID)
		if err != nil {
			return nil, err
		}

	}

	return deal, nil

}

func (r *StockDealRepository) GetAll(tp *models.TableParams, id int64) error {

	q := `
	SELECT sd.id, sd.stock_instrument_id, sd.strategy_id, sd.pattern_id, sd.currency, 
	sd.position, sd.time_frame, sd.enter_datetime, sd.enter_point, sd.stop_loss, 
	sd.quantity, sd.exit_datetime, sd.exit_point, sd.risk_ratio, sd.variability, 
	sd.user_id 
	FROM stock_deals AS sd
	WHERE sd.user_id=$1
	LIMIT $2 
	OFFSET $3;
	`

	rows, err := r.db.Query(
		q,
		id,
		tp.Pagination.ItemsPerPage,
		tp.Pagination.PageNumber*tp.Pagination.ItemsPerPage,
	)
	if err != nil {
		return err
	}

	source := []models.StockDeal{}

	count := 0

	for rows.Next() {
		count++
		var sd models.StockDeal
		err = rows.Scan(&sd.ID, &sd.Stock.ID, &sd.Strategy.ID, &sd.Pattern.ID,
			&sd.Currency, &sd.Position, &sd.TimeFrame, &sd.EnterDateTime,
			&sd.EnterPoint, &sd.StopLoss, &sd.Quantity, &sd.ExitDateTime,
			&sd.ExitPoint, &sd.RiskRatio, &sd.Variability, &sd.UserID)
		if err != nil {
			return err
		}

		source = append(source, sd)
	}

	if count == 0 {
		return nil
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
