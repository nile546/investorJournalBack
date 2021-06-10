package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/store"
)

type Repository struct {
	db                         *sql.DB
	userRepository             *UserRepository
	stockInstrumentRepository  *StockInstrumentRepository
	bankInstrumentRepository   *BankInstrumentRepository
	cryptoInstrumentRepository *CryptoInstrumentRepository
	stockDealRepository        *StockDealRepository
	stockStrategyRepository    *StockStrategyRepository
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

func (r *Repository) StockInstrument() store.StockInstrumentRepository {
	if r.stockInstrumentRepository != nil {
		return r.stockInstrumentRepository
	}

	r.stockInstrumentRepository = &StockInstrumentRepository{
		db: r.db,
	}

	return r.stockInstrumentRepository
}

func (r *Repository) BankInstrument() store.BankInstrumentRepository {
	if r.bankInstrumentRepository != nil {
		return r.bankInstrumentRepository
	}

	r.bankInstrumentRepository = &BankInstrumentRepository{
		db: r.db,
	}

	return r.bankInstrumentRepository
}

func (r *Repository) CryptoInstrument() store.CryptoInstrumentRepository {
	if r.cryptoInstrumentRepository != nil {
		return r.cryptoInstrumentRepository
	}

	r.cryptoInstrumentRepository = &CryptoInstrumentRepository{
		db: r.db,
	}

	return r.cryptoInstrumentRepository
}

func (r *Repository) StockDeal() store.StockDealRepository {
	if r.stockDealRepository != nil {
		return r.stockDealRepository
	}

	r.stockDealRepository = &StockDealRepository{
		db: r.db,
	}

	return r.stockDealRepository
}

func (r *Repository) StockStrategy() store.StockStrategyRepository {
	if r.stockStrategyRepository != nil {
		return r.stockStrategyRepository
	}

	r.stockStrategyRepository = &StockStrategyRepository{
		db: r.db,
	}

	return r.stockStrategyRepository
}
