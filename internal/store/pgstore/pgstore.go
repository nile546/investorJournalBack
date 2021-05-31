package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/store"
)

type Repository struct {
	db               *sql.DB
	userRepository   *UserRepository
	stockRepository  *StockRepository
	bankRepository   *BankRepository
	cryptoRepository *CryptoRepository
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

func (r *Repository) Bank() store.BankRepository {
	if r.bankRepository != nil {
		return r.bankRepository
	}

	r.bankRepository = &BankRepository{
		db: r.db,
	}

	return r.bankRepository
}

func (r *Repository) Crypto() store.CryptoRepository {
	if r.cryptoRepository != nil {
		return r.cryptoRepository
	}

	r.cryptoRepository = &CryptoRepository{
		db: r.db,
	}

	return r.cryptoRepository
}
