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

	bank := &models.BankInstrument{}

	deal := &models.DepositDeal{
		ID:   id,
		Bank: *bank,
	}

	err := r.db.QueryRow(q, id).Scan(
		&deal.Bank.ID,
		&deal.Currency,
		&deal.EnterDateTime,
		&deal.Percent,
		&deal.ExitDateTime,
		&deal.StartDeposit,
		&deal.EndDeposit,
		&deal.Result,
		&deal.UserID,
	)
	if err != nil {
		return nil, err
	}

	return deal, nil

}

func (r *DepositDealRepository) GetAll(tp *models.TableParams, id int64) error {

	q := `
	SELECT dd.id, dd.bank_instrument_id, dd.currency, dd.enter_datetime, dd.percent, 
	dd.exit_datetime, dd.start_deposit, dd.end_deposit, dd.reult, dd.user_id 
	FROM deposit_deals AS dd
	WHERE dd.user_id=$1
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

	source := []models.DepositDeal{}

	bank := &models.BankInstrument{}

	count := 0

	for rows.Next() {
		count++
		dd := models.DepositDeal{
			Bank: *bank,
		}
		err = rows.Scan(&dd.ID, dd.Bank.ID, dd.Currency,
			dd.EnterDateTime, dd.Percent, dd.ExitDateTime,
			dd.StartDeposit, dd.EndDeposit, dd.Result,
			dd.UserID)
		if err != nil {
			return err
		}

		source = append(source, dd)
	}

	if count == 0 {
		return nil
	}

	tp.Source = source
	return nil
}
