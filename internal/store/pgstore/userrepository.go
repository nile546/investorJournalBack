package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func (ur *UserRepository) CreateUser(u *models.User) error {

	q := `INSERT INTO users (login, email, encrypted_password) VALUES ($1, $2, $3) RETURNING id`

	if err := ur.db.QueryRow(q, u.Login, u.Email, u.EncryptedPassword).Scan(&u.ID); err != nil {
		return err
	}

	return nil

}

func (ur *UserRepository) UpdateUser(u *models.User) error {

	q := `UPDATE users SET (login, email, is_active) = ($1, $2, $3) WHERE id = $4`

	res, err := ur.db.Exec(q, u.Login, u.Email, u.IsActive, u.ID)
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

func (ur *UserRepository) UpdateIsActiveByUserID(ID int64) error {

	q := `UPDATE users SET is_active = $1 WHERE id = $2`

	res, err := ur.db.Exec(q, true, ID)
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

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {

	q := `SELECT id, login, encrypted_password, is_active, created_at FROM users where email=$1`

	res, err := ur.db.Query(q, email)
	if err != nil {
		return nil, err
	}

	u := &models.User{
		Email: email,
	}

	for res.Next() {

		err = res.Scan(&u.ID, &u.Login, &u.EncryptedPassword, &u.IsActive, &u.CreatedAt)
		if err != nil {
			return nil, err
		}

	}

	return u, nil
}

func (ur *UserRepository) GetUserByID(ID int64) (*models.User, error) {

	q := `SELECT email, login, encrypted_password, is_active, created_at FROM users where id=$1`

	res, err := ur.db.Query(q, ID)
	if err != nil {
		return nil, err
	}

	u := &models.User{
		ID: ID,
	}

	for res.Next() {

		err = res.Scan(&u.Email, &u.Login, &u.EncryptedPassword, &u.IsActive, &u.CreatedAt)
		if err != nil {
			return nil, err
		}

	}

	return u, nil
}

func (ur *UserRepository) SetRefreshToken(userID int64) (string, error) {

	q := `INSERT INTO refresh_tokens (user_id) VALUES ($1) RETURNING token`

	var token string

	if err := ur.db.QueryRow(q, userID).Scan(&token); err != nil {
		return "", err
	}

	return token, nil
}
