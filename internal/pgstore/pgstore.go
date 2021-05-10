package pgstore

import "database/sql"

type Repository struct {
	db *sql.DB
}

type UserRepository struct {
}

func New(db *sql.DB) *Repository {

	return &Repository{
		db: db,
	}
}

func (r *Repository) User() *UserRepository {

}
