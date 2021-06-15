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

	res, err := s.db.Exec(q, strategy.Name, strategy.Description, strategy.UserID, strategy.Type)
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

func (s *StrategyRepository) GetAllStrategy(userId int64) (*[]models.Strategy, error) {

	q := `SELECT * FROM strategies where user_id IS NULL OR user_id=$1`

	res, err := s.db.Query(q, userId)
	if err != nil {
		return nil, err
	}

	strgs := &[]models.Strategy{}

	for res.Next() {
		strg := models.Strategy{}
		err = res.Scan(&strg.ID, &strg.Name, &strg.Description, &strg.UserID, &strg.CreatedAt)
		if err != nil {
			return nil, err
		}
		*strgs = append(*strgs, strg)
	}

	return strgs, nil
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
