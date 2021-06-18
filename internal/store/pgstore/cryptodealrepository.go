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

	res, err := r.db.Exec(q, deal.Crypto.ID, deal.Currency, deal.Strategy.ID,
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

	crypto := &models.CryptoInstrument{}
	strategy := &models.Strategy{}
	pattern := &models.Pattern{}

	deal := &models.CryptoDeal{
		ID:       id,
		Crypto:   *crypto,
		Strategy: strategy,
		Pattern:  pattern,
	}

	err := r.db.QueryRow(q, id).Scan(
		&deal.Crypto.ID,
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
		&deal.RiskRatio,
		&deal.UserID,
	)
	if err != nil {
		return nil, err
	}

	return deal, nil

}

func (r *CryptoDealRepository) GetAll(tp *models.TableParams, id int64) error {

	q := `
	SELECT cd.id, cd.crypto_instrument_id, cd.currency, cd.strategy_id,
	cd.pattern_id, cd.position, cd.time_frame, cd.enter_date_time, cd.enter_point, 
	cd.stop_loss, cd.quantity, cd.exit_datetime, cd.exit_point, cd.risk_ratio, cd.user_id
	FROM crypto_deals AS cd
	WHERE cd.user_id=$1
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

	count := 0

	crypto := &models.CryptoInstrument{}
	strategy := &models.Strategy{}
	pattern := &models.Pattern{}

	source := []models.CryptoDeal{}

	for rows.Next() {
		count++
		cd := models.CryptoDeal{
			ID:       id,
			Crypto:   *crypto,
			Strategy: strategy,
			Pattern:  pattern,
		}
		err = rows.Scan(&cd.ID, &cd.Crypto.ID, &cd.Strategy.ID, &cd.Pattern,
			&cd.Currency, &cd.Position, &cd.TimeFrame, &cd.EnterDateTime,
			&cd.EnterPoint, &cd.StopLoss, &cd.Quantity, &cd.ExitDateTime,
			&cd.ExitPoint, &cd.RiskRatio, &cd.UserID)
		if err != nil {
			return err
		}

		source = append(source, cd)
	}

	if count == 0 {
		return nil
	}

	tp.Source = source
	return nil
}
