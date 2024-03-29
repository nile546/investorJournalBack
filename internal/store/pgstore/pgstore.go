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
	strategyRepository         *StrategyRepository
	patternRepository          *PatternRepository
	stockDealPartRepository    *StockDealPartRepository
	tinkoffTokenRepository     *TinkoffTokenRepository
	cryptoDealRepository       *CryptoDealRepository
	depositDealRepository      *DepositDealRepository
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

func (r *Repository) Strategy() store.StrategyRepository {
	if r.strategyRepository != nil {
		return r.strategyRepository
	}

	r.strategyRepository = &StrategyRepository{
		db: r.db,
	}

	return r.strategyRepository
}

func (r *Repository) Pattern() store.PatternRepository {
	if r.patternRepository != nil {
		return r.patternRepository
	}

	r.patternRepository = &PatternRepository{
		db: r.db,
	}

	return r.patternRepository
}

func (r *Repository) StockDealPart() store.StockDealPartRepository {
	if r.stockDealPartRepository != nil {
		return r.stockDealPartRepository
	}

	r.stockDealPartRepository = &StockDealPartRepository{
		db: r.db,
	}

	return r.stockDealPartRepository
}

func (r *Repository) TinkoffToken() store.TinkoffTokenRepository {
	if r.tinkoffTokenRepository != nil {
		return r.tinkoffTokenRepository
	}

	r.tinkoffTokenRepository = &TinkoffTokenRepository{
		db: r.db,
	}

	return r.tinkoffTokenRepository
}

func (r *Repository) CryptoDeal() store.CryptoDealRepository {
	if r.cryptoDealRepository != nil {
		return r.cryptoDealRepository
	}

	r.cryptoDealRepository = &CryptoDealRepository{
		db: r.db,
	}

	return r.cryptoDealRepository
}

func (r *Repository) DepositDeal() store.DepositDealRepository {
	if r.depositDealRepository != nil {
		return r.depositDealRepository
	}

	r.depositDealRepository = &DepositDealRepository{
		db: r.db,
	}

	return r.depositDealRepository
}
