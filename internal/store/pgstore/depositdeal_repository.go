package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type DepositDealRepository struct {
	db *sql.DB
}

func (r *DepositDealRepository) CreateDepositDeal(deal *models.DepositDeal) error {

	q := `INSERT INTO deposit_deals
	(bank_instrument_id, currency, enter_datetime, percent, 
	exit_datetime, start_deposit, end_deposit, result, user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	res, err := r.db.Exec(q, deal.Bank.ID, deal.Currency,
		deal.EnterDateTime, deal.Percent, deal.ExitDateTime,
		deal.StartDeposit, deal.EndDeposit, deal.Result,
		deal.UserID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *DepositDealRepository) UpdateDepositDeal(deal *models.DepositDeal) error {

	q := `UPDATE deposit_deals SET (bank_instrument_id, currency, enter_datetime, percent, 
		exit_datetime, start_deposit, end_deposit, result, user_id)=($1, $2, $3, $4, $5, $6,
		$7, $8, $9, $10) WHERE id=$11`

	res, err := r.db.Exec(q, deal.Bank.ID, deal.Currency,
		deal.EnterDateTime, deal.Percent, deal.ExitDateTime,
		deal.StartDeposit, deal.EndDeposit, deal.Result,
		deal.UserID, deal.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *DepositDealRepository) DeleteDepositDeal(id int64) error {

	q := "DELETE FROM deposit_deals WHERE id=$1"

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

func (r *DepositDealRepository) GetDepositDealByID(id int64) (*models.DepositDeal, error) {

	q := `SELECT bank_instrument_id, currency, enter_datetime, percent, 
	exit_datetime, start_deposit, end_deposit, result, user_id FROM deposit_deals where id=$1`

	res, err := r.db.Query(q, id)
	if err != nil {
		return nil, err
	}

	deal := &models.DepositDeal{
		ID: id,
	}

	for res.Next() {

		err = res.Scan(deal.Bank.ID, deal.Currency,
			deal.EnterDateTime, deal.Percent, deal.ExitDateTime,
			deal.StartDeposit, deal.EndDeposit, deal.Result,
			deal.UserID, deal.ID)
		if err != nil {
			return nil, err
		}

	}

	return deal, nil

}

func (r *DepositDealRepository) GetAll(tp *models.TableParams) error {

	q := `
	SELECT sd.id 
	FROM deposit_deals AS sd
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

	source := []models.DepositDeal{}

	for rows.Next() {
		var sd models.DepositDeal
		err = rows.Scan(&sd.ID)
		if err != nil {
			return err
		}

		source = append(source, sd)
	}

	tp.Source = source
	return nil
}
