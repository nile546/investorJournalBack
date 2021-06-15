package store

import "github.com/nile546/diplom/internal/models"

type Repository interface {
	User() UserRepository
	StockInstrument() StockInstrumentRepository
	BankInstrument() BankInstrumentRepository
	CryptoInstrument() CryptoInstrumentRepository
	StockDeal() StockDealRepository
	Strategy() StrategyRepository
	Pattern() PatternRepository
}

type UserRepository interface {
	CreateUser(*models.User) error
	UpdateUser(*models.User) error
	UpdateIsActiveByUserID(ID int64) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(ID int64) (*models.User, error)
	SetRefreshToken(int64) (string, error)
	UpdateRefreshToken(string) (string, int64, error)
	DeleteRefreshTokenByUser(*models.User) error
}

type StockInstrumentRepository interface {
	InsertStocksInstruments(*[]models.StockInstrument) error
	GetAllStockInstruments() (*[]models.StockInstrument, error)
}

type BankInstrumentRepository interface {
	InsertBanksInstruments(*[]models.BankInstrument) error
	GetAllBankInstruments() (*[]models.BankInstrument, error)
}

type CryptoInstrumentRepository interface {
	InsertCryptoInstruments(*[]models.CryptoInstrument) error
	GetAllCryptoInstruments() (*[]models.CryptoInstrument, error)
}

type StockDealRepository interface {
	GetAll(*models.TableParams) error
}

type StrategyRepository interface {
	CreateStrategy(*models.Strategy) error
	UpdateStrategy(*models.Strategy) error
	GetAllStrategy(int64) (*[]models.Strategy, error)
	DeleteStrategy(int64) error
}

type PatternRepository interface {
	CreatePattern(*models.Pattern) error
	UpdatePattern(*models.Pattern) error
	GetAllPattern(int64) (*[]models.Pattern, error)
	DeletePattern(int64) error
}
