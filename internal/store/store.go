package store

import (
	"time"

	"github.com/nile546/diplom/internal/models"
)

type Repository interface {
	User() UserRepository
	StockInstrument() StockInstrumentRepository
	BankInstrument() BankInstrumentRepository
	CryptoInstrument() CryptoInstrumentRepository
	StockDeal() StockDealRepository
	Strategy() StrategyRepository
	Pattern() PatternRepository
	StockDealPart() StockDealPartRepository
	TinkoffToken() TinkoffTokenRepository
	CryptoDeal() CryptoDealRepository
	DepositDeal() DepositDealRepository
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
	UpdateDateGrab(time.Time, int64) error
	GetDateGrabByUserID(int64) (time.Time, error)
	UpdateAutoGrab(int64, bool) error
}

type StockInstrumentRepository interface {
	InsertStocksInstruments(*[]models.StockInstrument) error
	GetAllStockInstruments() (*[]models.StockInstrument, error)
	GetInstrumentByISIN(string) (*models.StockInstrument, error)
	GetPopularStockInstrumentByUserID(int64) (*models.StockInstrument, error)
	GetPopularStockInstrumentsID() ([]int64, error)
	GetStockInstrumentByID(int64) (*models.StockInstrument, error)
	GetAll(*models.TableParams) error
}

type BankInstrumentRepository interface {
	InsertBanksInstruments(*[]models.BankInstrument) error
	GetAllBankInstruments() (*[]models.BankInstrument, error)
	GetPopularBankInstrumentByUserID(int64) (*models.BankInstrument, error)
	GetPopularBankInstrumentsID() ([]int64, error)
	GetBankInstrumentByID(int64) (*models.BankInstrument, error)
	GetAll(*models.TableParams) error
}

type CryptoInstrumentRepository interface {
	InsertCryptoInstruments(*[]models.CryptoInstrument) error
	GetAllCryptoInstruments() (*[]models.CryptoInstrument, error)
	GetPopularCryptoInstrumentByUserID(int64) (*models.CryptoInstrument, error)
	GetPopularCryptoInstrumentsID() ([]int64, error)
	GetCryptoInstrumentByID(int64) (*models.CryptoInstrument, error)
	GetAll(*models.TableParams) error
}

type StockDealRepository interface {
	CreateStockDeal(*models.StockDeal) error
	UpdateStockDeal(*models.StockDeal) error
	DeleteStockDeal(int64) error
	GetStockDealByID(int64) (*models.StockDeal, error)
	GetAll(*models.TableParams, int64) error
	GetStockDealsIDByISIN(string) int64
	CreateOpenStockDeal(*models.StockDeal) (int64, error)
	UpdateQuantityStockDeal(int64, int) error
	SetStockDealCompleted(time.Time, int64, int64) error
}

type StrategyRepository interface {
	CreateStrategy(*models.Strategy) error
	UpdateStrategy(*models.Strategy) error
	GetAllStrategy(*models.TableParams, int64) error
	DeleteStrategy(int64) error
}

type PatternRepository interface {
	CreatePattern(*models.Pattern) error
	UpdatePattern(*models.Pattern) error
	GetAllPattern(*models.TableParams, int64) error
	DeletePattern(int64) error
}

type StockDealPartRepository interface {
	InsertStockDealPart(*models.StockDealParts) error
	CheckQuantityDeal(int64) (bool, error)
}

type TinkoffTokenRepository interface {
	InsertTinkoffToken(string, int64) error
	GetTinkoffToken(int64) (string, error)
	UpdateTinkoffToken(string, int64) error
}

type CryptoDealRepository interface {
	CreateCryptoDeal(*models.CryptoDeal) error
	UpdateCryptoDeal(*models.CryptoDeal) error
	DeleteCryptoDeal(int64) error
	GetCryptoDealByID(int64) (*models.CryptoDeal, error)
	GetAll(*models.TableParams, int64) error
}

type DepositDealRepository interface {
	CreateDepositDeal(*models.DepositDeal) error
	UpdateDepositDeal(*models.DepositDeal) error
	DeleteDepositDeal(int64) error
	GetDepositDealByID(int64) (*models.DepositDeal, error)
	GetAll(*models.TableParams, int64) error
}
