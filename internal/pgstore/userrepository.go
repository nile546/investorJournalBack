package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func (ur *UserRepository) Create(u *models.User) error {

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.EncryptPass(); err != nil {
		return err
	}

	q := `INSERT INTO users (name, login, encrypted_password) VALUES ($1, $2, $3) RETURNING id`

	return ur.db.QueryRow(q, u.Login, u.Email, u.EncryptedPassword).Scan(&u.ID)
}