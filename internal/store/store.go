package store

import "github.com/nile546/diplom/internal/models"

type Repository interface {
	User() UserRepository
	StockInstrument() StockInstrumentRepository
	BankInstrument() BankInstrumentRepository
	CryptoInstrument() CryptoInstrumentRepository
	StockDeal() StockDealRepository
	StockStrategy() StockStrategyRepository
}

type UserRepository interface {
	CreateUser(*models.User) error
	UpdateUser(*models.User) error
	UpdateIsActiveByUserID(ID int64) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(ID int64) (*models.User, error)
}

type StockInstrumentRepository interface {
	InsertStocksInstruments(*[]models.StockInstrument) error
	TruncateStocksInstruments() error
}

type BankInstrumentRepository interface {
	InsertBanksInstruments(*[]models.BankInstrument) error
	TruncateBanksInstruments() error
}

type CryptoInstrumentRepository interface {
	InsertCryptoInstruments(*[]models.CryptoInstrument) error
	TruncateCryptoInstruments() error
}

type StockDealRepository interface {
}

type StockStrategyRepository interface {
	CreateStockStrategy(*models.StockStrategy) error
	UpdateStockStrategy(*models.StockStrategy) error
	GetAllStockStrategy(*int64) (*[]models.StockStrategy, error)
	DeleteStockStrategy(*int64) error
}
