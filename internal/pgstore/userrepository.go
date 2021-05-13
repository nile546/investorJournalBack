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

	q := `INSERT INTO users (login, email, encrypted_password) VALUES ($1, $2, $3) RETURNING id`

	if err := ur.db.QueryRow(q, u.Login, u.Email, u.EncryptedPassword).Scan(&u.ID); err != nil {
		return err
	}

	return nil

}

func (ur *UserRepository) Update(u *models.User) error {

	q := `UPDATE users SET (login, email, is_active, registration_token) = ($1, $2, $3, $4) WHERE id = $5`

	res, err := ur.db.Exec(q, u.Login, u.Email, u.IsActive, u.RegistrationToken, u.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count < 1 {
		//Add err consts
		return nil
	}

	return nil
}

func (ur *UserRepository) Check(email string, password string) error {

	//pass encrypt

	q := `SELECT encrypted_password FROM users WHERE email = $1`

	res, err := ur.db.Exec(q, email)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count < 1 {
		//Add err consts
		return nil
	}

	return nil
}
