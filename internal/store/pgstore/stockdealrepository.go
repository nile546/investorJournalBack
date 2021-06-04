package pgstore

import "database/sql"

type StockDealRepository struct {
	db *sql.DB
}
