package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type StockPatternRepository struct {
	db *sql.DB
}

func (s *StockPatternRepository) CreateStockPattern(pattern *models.StockPattern) error {

	q := "INSERT INTO stock_patterns (name, description, user_id) VALUES ($1, $2, $3)"

	res, err := s.db.Exec(q, pattern.Name, pattern.Description, pattern.UserID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *StockPatternRepository) UpdateStockPattern(pattern *models.StockPattern) error {

	q := "UPDATE stock_patterns SET (name, description) = ($1, $2) WHERE id = $3"

	res, err := s.db.Exec(q, pattern.Name, pattern.Description, pattern.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *StockPatternRepository) GetAllStockPattern(userId int64) (*[]models.StockPattern, error) {

	q := `SELECT * FROM stock_patterns where user_id=$1`

	res, err := s.db.Query(q, userId)
	if err != nil {
		return nil, err
	}

	ptrns := &[]models.StockPattern{}

	for res.Next() {
		ptrn := models.StockPattern{}
		err = res.Scan(&ptrn.ID, &ptrn.Name, &ptrn.Description, &ptrn.Icon, &ptrn.UserID, &ptrn.CreatedAt)
		if err != nil {
			return nil, err
		}
		*ptrns = append(*ptrns, ptrn)
	}

	return ptrns, nil
}

func (s *StockPatternRepository) DeleteStockPattern(id int64) error {

	q := "DELETE FROM stock_patterns WHERE id=$1"

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
