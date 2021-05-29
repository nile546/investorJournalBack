package store

import "github.com/nile546/diplom/internal/models"

type Repository interface {
	User() UserRepository
	Stock() StockRepository
	Bank() DepositRepository
}

type UserRepository interface {
	Create(*models.User) error
	Update(*models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(ID int64) (*models.User, error)
}

type StockRepository interface {
	InsertStocks(*[]models.Stock) error
}

type DepositRepository interface {
	InsertBanks(*[]models.Bank) error
}
