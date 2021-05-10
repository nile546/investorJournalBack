package store

import "github.com/nile546/diplom/internal/models"

type Repository interface {
	User() UserRepository
}

type UserRepository interface {
	Create(*models.User) error
}
