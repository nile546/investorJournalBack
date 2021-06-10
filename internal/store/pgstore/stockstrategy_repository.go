package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type StockStrategyRepository struct {
	db *sql.DB
}

func (s *StockStrategyRepository) CreateStockStrategy(strategy *models.StockStrategy) error {

	q := "INSERT INTO stock_strategy (name, description, user_id) VALUES ($1, $2, $3)"

	res, err := s.db.Exec(q, strategy.Name, strategy.Description, strategy.UserID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *StockStrategyRepository) UpdateStockStrategy(strategy *models.StockStrategy) error {

	q := "UPDATE stock_strategy SET (name, description) = ($1, $2) WHERE id = $3"

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

func (s *StockStrategyRepository) GetAllStockStrategy(userId *int64) (strgs *[]models.StockStrategy, err error) {

	q := `SELECT * FROM stock_strategy where user_id=$1`

	res, err := s.db.Query(q, userId)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		strg := models.StockStrategy{}
		err = res.Scan(&strg.ID, &strg.Name, &strg.Description, &strg.UserID, &strg.CreatedAt)
		if err != nil {
			return nil, err
		}
		*strgs = append(*strgs, strg)
	}

	return strgs, nil
}

func (s *StockStrategyRepository) DeleteStockStrategy(id *int64) error {

	q := "DELETE FROM stock_strategy WHERE id=$1"

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
