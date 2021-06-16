package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type StockDealPartRepository struct {
	db *sql.DB
}

func (s *StockDealPartRepository) InsertStockDealPart(part *models.StockDealParts) error {

	q := "INSERT INTO stockdeal_parts(quantity, price, deal_type, datetime, stock_deal_id) VALUES ($1, $2, $3, $4, $5)"

	res, err := s.db.Exec(q, part.Quantity, part.Price, part.DealType, part.DateTime, part.StockDealId)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *StockDealPartRepository) CheckQuantityDeal(idStockDeal int64) (bool, error) {

	q := `SELECT
	CASE
		WHEN (SELECT SUM(quantity) FROM stockdeal_parts WHERE stock_deal_id = $1 AND type = 1)
		   = (SELECT SUM(quantity) FROM stockdeal_parts WHERE stock_deal_id = $1 AND type = 2)  THEN true
		ELSE false
	END as check_quantity`

	var statusDeal bool

	if err := s.db.QueryRow(q, idStockDeal).Scan(&statusDeal); err != nil {
		return false, err
	}

	return statusDeal, nil
}
