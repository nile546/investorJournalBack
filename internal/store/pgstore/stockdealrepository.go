package pgstore

import (
	"database/sql"
	"math"
	"time"

	"github.com/nile546/diplom/internal/models"
)

type StockDealRepository struct {
	db *sql.DB
}

func (r *StockDealRepository) CreateStockDeal(stockDeal *models.StockDeal, userId int64) error {

	q := `INSERT INTO stock_deals
	(
		stock_instrument_id, 
		currency,
		strategy_id, 
		pattern_id, 
		position, 
		time_frame, 
		enter_datetime, 
		enter_point, 
		stop_loss, 
		quantity, 
		exit_datetime, 
		exit_point, 
		user_id
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	res, err := r.db.Exec(
		q,
		stockDeal.Stock.ID,
		stockDeal.Currency,
		stockDeal.Strategy.ID,
		stockDeal.Pattern.ID,
		stockDeal.Position,
		stockDeal.TimeFrame,
		stockDeal.EnterDateTime,
		stockDeal.EnterPoint,
		stockDeal.StopLoss,
		stockDeal.Quantity,
		stockDeal.ExitDateTime,
		stockDeal.ExitPoint,
		userId,
	)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *StockDealRepository) UpdateStockDeal(stockDeal *models.StockDeal, userId int64) error {

	q := `UPDATE stock_deals 
	SET 
	(
		stock_instrument_id,
		currency,
		strategy_id,
		pattern_id, 
		position,
		time_frame,
		enter_datetime,
		enter_point,
		stop_loss, 
		quantity,
		exit_datetime,
		exit_point,
	)=($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	WHERE id=$14`

	res, err := r.db.Exec(
		q,
		stockDeal.Stock.ID,
		stockDeal.Currency,
		stockDeal.Strategy.ID,
		stockDeal.Pattern.ID,
		stockDeal.Position,
		stockDeal.TimeFrame,
		stockDeal.EnterDateTime,
		stockDeal.EnterPoint,
		stockDeal.StopLoss,
		stockDeal.Quantity,
		stockDeal.ExitDateTime,
		stockDeal.ExitPoint,
		stockDeal.ID,
	)
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

	q := `SELECT 
	stock_instrument_id,
	strategy_id,
	pattern_id,
	currency, 
	position,
	time_frame,
	enter_datetime,
	enter_point,
	stop_loss, 
	quantity,
	exit_datetime,
	exit_point,
	variability, 
	FROM stock_deals where id=$1`
	stock := &models.StockInstrument{}
	strategy := &models.Strategy{}
	pattern := &models.Pattern{}

	deal := models.StockDeal{
		ID:       id,
		Stock:    *stock,
		Strategy: strategy,
		Pattern:  pattern,
	}

	err := r.db.QueryRow(q, id).Scan(
		&deal.Stock.ID,
		&deal.Strategy.ID,
		&deal.Pattern.ID,
		&deal.Currency,
		&deal.Position,
		&deal.TimeFrame,
		&deal.EnterDateTime,
		&deal.EnterPoint,
		&deal.StopLoss,
		&deal.Quantity,
		&deal.ExitDateTime,
		&deal.ExitPoint,
		&deal.Variability,
	)
	if err != nil {
		return nil, err
	}

	return &deal, nil

}

func (r *StockDealRepository) GetAll(tp *models.TableParams, userId int64) (*[]*models.StockDeal, error) {

	source := &[]*models.StockDeal{}

	q := `
	SELECT 
	sd.id,
	sd.stock_instrument_id,
    sd.currency,
    sd.strategy_id,
    sd.pattern_id,
    sd.position,
    sd.time_frame,
    sd.enter_datetime,
    sd.enter_point,
    sd.stop_loss,
    sd.quantity,
    sd.exit_datetime,
    sd.exit_point,
    sd.variability,

	st.title,
	st.ticker,

	s.name,

	p.name

	FROM stock_deals AS sd
	LEFT JOIN stocks_instruments AS st ON sd.stock_instrument_id = st.id
	LEFT JOIN strategies AS s ON sd.strategy_id = s.id
	LEFT JOIN patterns AS p ON sd.pattern_id = p.id
	WHERE sd.user_id = $1
	LIMIT $2
	OFFSET $3
	`

	rows, err := r.db.Query(
		q,
		userId,
		tp.Pagination.ItemsPerPage,
		tp.Pagination.PageNumber*tp.Pagination.ItemsPerPage,
	)
	if err != nil {
		return nil, err
	}

	stock := &models.StockInstrument{}
	strategy := &models.Strategy{}
	pattern := &models.Pattern{}

	count := 0

	for rows.Next() {
		count++
		sd := &models.StockDeal{
			Stock:    *stock,
			Strategy: strategy,
			Pattern:  pattern,
		}
		err = rows.Scan(
			&sd.ID,
			&sd.Stock.ID,
			&sd.Strategy.ID,
			&sd.Pattern.ID,
			&sd.Currency,
			&sd.Position,
			&sd.TimeFrame,
			&sd.EnterDateTime,
			&sd.EnterPoint,
			&sd.StopLoss,
			&sd.Quantity,
			&sd.ExitDateTime,
			&sd.ExitPoint,
			&sd.Variability,

			&sd.Stock.Title,
			&sd.Stock.Ticker,

			&sd.Strategy.Name,

			&sd.Pattern.Name,
		)
		if err != nil {
			return nil, err
		}

		*source = append(*source, sd)
	}

	if count == 0 {
		return nil, nil
	}

	q = `SELECT COUNT(id)
	FROM stock_deals
	WHERE user_id = $1
	`
	var itemsCount int
	if err = r.db.QueryRow(q, userId).Scan(&itemsCount); err != nil {
		return nil, err
	}
	defer rows.Close()

	tp.Pagination.PageCount = 0

	if itemsCount > 0 {
		tp.Pagination.PageCount = int(math.Ceil(float64(itemsCount) / float64(tp.Pagination.ItemsPerPage)))
	}

	return source, nil
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

func (r *StockDealRepository) CreateOpenStockDeal(stockDeal *models.StockDeal, userId int64) (int64, error) {

	q := `INSERT INTO stock_deals 
	(
		stock_instrument_id,
		currency,
		enter_datetime,
		enter_point,
		quantity,
		variability,
		user_id
	) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) 
	RETURNING id`

	var id int64

	err := r.db.QueryRow(
		q,
		stockDeal.Stock.ID,
		&stockDeal.Currency,
		stockDeal.EnterDateTime,
		stockDeal.EnterPoint,
		stockDeal.Quantity,
		stockDeal.Variability,
		userId,
	).Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
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

func (r *StockDealRepository) GetVariabilityByID(id int64) (bool, error) {

	q := `SELECT variability FROM stock_deals where id=$1`

	var variability bool

	err := r.db.QueryRow(q, id).Scan(
		&variability,
	)
	if err != nil {
		return false, err
	}

	return variability, nil

}
