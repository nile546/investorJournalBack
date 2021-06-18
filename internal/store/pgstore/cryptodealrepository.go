package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type CryptoDealRepository struct {
	db *sql.DB
}

func (r *CryptoDealRepository) CreateCryptoDeal(deal *models.CryptoDeal) error {

	q := `INSERT INTO crypto_deals
	(crypto_instrument_id, currency, strategy_id, pattern_id, 
	position, time_frame, enter_datetime, enter_point, stop_loss, 
	quantity, exit_datetime, exit_point, risk_ratio, user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	res, err := r.db.Exec(q, deal.Crypto, deal.Currency, deal.Strategy.ID,
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

func (r *CryptoDealRepository) UpdateCryptoDeal(deal *models.CryptoDeal) error {

	q := `UPDATE crypto_deals SET (crypto_instrument_id, currency, strategy_id, pattern_id, 
		position, time_frame, enter_datetime, enter_point, stop_loss, 
		quantity, exit_datetime, exit_point, risk_ratio)=($1, $2, $3, $4, $5, $6,
		$7, $8, $9, $10, $11, $12, $13)	WHERE id=$14`

	res, err := r.db.Exec(q, deal.Crypto.ID, deal.Currency, deal.Strategy.ID,
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

func (r *CryptoDealRepository) DeleteCryptoDeal(id int64) error {

	q := "DELETE FROM crypto_deals WHERE id=$1"

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

func (r *CryptoDealRepository) GetCryptoDealByID(id int64) (*models.CryptoDeal, error) {

	q := `SELECT crypto_instrument_id, strategy_id, pattern_id, currency, position, time_frame, enter_datetime, enter_point, stop_loss, 
	quantity, exit_datetime, exit_point, risk_ratio, user_id FROM crypto_deals where id=$1`

	res, err := r.db.Query(q, id)
	if err != nil {
		return nil, err
	}

	deal := &models.CryptoDeal{
		ID: id,
	}

	for res.Next() {

		err = res.Scan(&deal.Crypto.ID, &deal.Strategy.ID, &deal.Pattern,
			&deal.Currency, &deal.Position, &deal.TimeFrame, &deal.EnterDateTime,
			&deal.EnterPoint, &deal.StopLoss, &deal.Quantity, &deal.ExitDateTime,
			&deal.ExitPoint, &deal.RiskRatio, &deal.UserID)
		if err != nil {
			return nil, err
		}

	}

	return deal, nil

}

func (r *CryptoDealRepository) GetAll(tp *models.TableParams) error {

	q := `
	SELECT sd.id 
	FROM crypto_deals AS sd
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

	source := []models.CryptoDeal{}

	for rows.Next() {
		var sd models.CryptoDeal
		err = rows.Scan(&sd.ID)
		if err != nil {
			return err
		}

		source = append(source, sd)
	}

	tp.Source = source
	return nil
}
