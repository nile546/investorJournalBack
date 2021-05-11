package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/config"
	"github.com/nile546/diplom/internal/store"
)

type Repository struct {
	db             *sql.DB
	c              *config.Config
	userRepository *UserRepository
}

func New(db *sql.DB, c *config.Config) *Repository {
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
		c:  r.c,
	}

	return r.userRepository
}
