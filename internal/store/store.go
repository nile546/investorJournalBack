package store

import "github.com/nile546/diplom/internal/models"

type Repository interface {
	User() UserRepository
}

type UserRepository interface {
	Create(*models.User) error
	Update(*models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(ID int) (*models.User, error)
}
