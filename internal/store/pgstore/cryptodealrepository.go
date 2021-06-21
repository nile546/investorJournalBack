package pgstore

import (
	"database/sql"
	"math"

	"github.com/nile546/diplom/internal/models"
)

type CryptoDealRepository struct {
	db *sql.DB
}

func (r *CryptoDealRepository) CreateCryptoDeal(cd *models.CryptoDeal, userId int64) error {

	q := `INSERT INTO crypto_deals
	(
		crypto_instrument_id, 
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
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	res, err := r.db.Exec(
		q,
		cd.Crypto.ID,
		cd.Currency,
		cd.Strategy.GetID(),
		cd.Pattern.GetID(),
		cd.Position,
		cd.TimeFrame,
		cd.EnterDateTime,
		cd.EnterPoint,
		cd.StopLoss,
		cd.Quantity,
		cd.ExitDateTime,
		cd.ExitPoint,
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

func (r *CryptoDealRepository) UpdateCryptoDeal(cd *models.CryptoDeal, userId int64) error {

	q := `UPDATE crypto_deals 
	SET (
		crypto_instrument_id,
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
		exit_point
		)=($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		WHERE id = $13`

	res, err := r.db.Exec(
		q,
		cd.Crypto.ID,
		cd.Currency,
		cd.Strategy.ID,
		cd.Pattern.ID,
		cd.Position,
		cd.TimeFrame,
		cd.EnterDateTime,
		cd.EnterPoint,
		cd.StopLoss,
		cd.Quantity,
		cd.ExitDateTime,
		cd.ExitPoint,
		cd.ID,
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

	q := `SELECT 
	crypto_instrument_id,
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
	FROM crypto_deals where id = $1`

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
	)
	if err != nil {
		return nil, err
	}

	return deal, nil

}

func (r *CryptoDealRepository) GetAll(tp *models.TableParams, userId int64) (*[]*models.CryptoDeal, error) {

	source := &[]*models.CryptoDeal{}

	q := `
	SELECT 
	cd.id,
	cd.crypto_instrument_id,
	cd.currency,
	cd.strategy_id,
	cd.pattern_id,
	cd.position,
	cd.time_frame,
	cd.enter_datetime,
	cd.enter_point, 
	cd.stop_loss,
	cd.quantity,
	cd.exit_datetime,
	cd.exit_point,

	cr.title,
	cr.ticker,

	s.name,

	p.name

	FROM crypto_deals AS cd
	LEFT JOIN crypto_instruments AS cr ON cd.crypto_instrument_id = cr.id
	LEFT JOIN strategies AS s ON cd.strategy_id = s.id
	LEFT JOIN patterns AS p ON cd.pattern_id = p.id
	WHERE cd.user_id = $1
	LIMIT $2 
	OFFSET $3;
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

	count := 0

	for rows.Next() {
		count++

		crypto := &models.CryptoInstrument{}

		var strategyID *int64
		var strategyName *string

		var patternID *int64
		var patternName *string

		cd := &models.CryptoDeal{
			Crypto: *crypto,
		}

		err = rows.Scan(
			&cd.ID,
			&cd.Crypto.ID,
			&strategyID,
			&patternID,
			&cd.Currency,
			&cd.Position,
			&cd.TimeFrame,
			&cd.EnterDateTime,
			&cd.EnterPoint,
			&cd.StopLoss,
			&cd.Quantity,
			&cd.ExitDateTime,
			&cd.ExitPoint,

			&cd.Crypto.Title,
			&cd.Crypto.Ticker,

			&strategyName,

			&patternName,
		)
		if err != nil {
			return nil, err
		}

		if strategyID != nil {
			cd.Strategy = &models.Strategy{
				ID:   *strategyID,
				Name: strategyName,
			}
		}

		if patternID != nil {
			cd.Pattern = &models.Pattern{
				ID:   *patternID,
				Name: patternName,
			}
		}

		*source = append(*source, cd)
	}

	if count == 0 {
		return nil, nil
	}

	q = `SELECT COUNT(id)
	FROM crypto_deals
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
