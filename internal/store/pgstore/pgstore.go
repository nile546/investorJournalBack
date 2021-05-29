package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/store"
)

type Repository struct {
	db                *sql.DB
	userRepository    *UserRepository
	stockRepository   *StockRepository
	depositRepository *DepositRepository
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) User() store.UserRepository {

	if r.userRepository != nil {
		return r.userRepository
	}

	r.userRepository = &UserRepository{
		db: r.db,
	}

	return r.userRepository
}

func (r *Repository) Stock() store.StockRepository {
	if r.stockRepository != nil {
		return r.stockRepository
	}

	r.stockRepository = &StockRepository{
		db: r.db,
	}

	return r.stockRepository
}

func (r *Repository) Bank() store.DepositRepository {
	if r.depositRepository != nil {
		return r.depositRepository
	}

	r.depositRepository = &DepositRepository{
		db: r.db,
	}

	return r.depositRepository
}
