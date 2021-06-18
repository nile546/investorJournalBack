package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type StrategyRepository struct {
	db *sql.DB
}

func (s *StrategyRepository) CreateStrategy(strategy *models.Strategy) error {

	q := "INSERT INTO strategies (name, description, user_id, type) VALUES ($1, $2, $3, $4)"

	res, err := s.db.Exec(q, strategy.Name, strategy.Description, strategy.UserID, strategy.InstrumentType)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *StrategyRepository) UpdateStrategy(strategy *models.Strategy) error {

	q := "UPDATE strategies SET (name, description) = ($1, $2) WHERE id = $3"

	res, err := s.db.Exec(q, strategy.Name, strategy.Description, strategy.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *StrategyRepository) GetAllStrategy(tp *models.TableParams, id int64) error {

	//Add user_id IS NULL OR user_id=$1
	q := `SELECT id, name, description, user_id, instrument_type, created_at
	FROM strategies
	WHERE user_id IS NULL OR user_id=$1
	LIMIT $2
	OFFSET $3;
	`

	rows, err := s.db.Query(
		q,
		id,
		tp.Pagination.ItemsPerPage,
		tp.Pagination.PageNumber*tp.Pagination.ItemsPerPage,
	)
	if err != nil {
		return err
	}

	source := []models.Strategy{}

	count := 0

	for rows.Next() {
		count++
		var s models.Strategy
		err = rows.Scan(&s.ID, &s.Name, &s.Description, &s.UserID, &s.InstrumentType, &s.CreatedAt)
		if err != nil {
			return err
		}

		source = append(source, s)
	}

	if count == 0 {
		return nil
	}

	tp.Source = source

	return nil
}

func (s *StrategyRepository) DeleteStrategy(id int64) error {

	q := "DELETE FROM strategies WHERE id=$1"

	res, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil

}
