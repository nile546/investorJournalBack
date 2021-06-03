package store

import "github.com/nile546/diplom/internal/models"

type Repository interface {
	User() UserRepository
	Stock() StockRepository
	Bank() BankRepository
	Crypto() CryptoRepository
}

type UserRepository interface {
	Create(*models.User) error
	Update(*models.User) error
	UpdateIsActiveByUserID(ID int64) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(ID int64) (*models.User, error)
}

type StockRepository interface {
	InsertStocks(*[]models.Stock) error
	TruncateStocks() error
}

type BankRepository interface {
	InsertBanks(*[]models.Bank) error
	TruncateBanks() error
}

type CryptoRepository interface {
	InsertCrypto(*[]models.Crypto) error
	TruncateCrypto() error
}
