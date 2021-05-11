package pgstore

import (
	"database/sql"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nile546/diplom/config"
	"github.com/nile546/diplom/internal/models"
)

type UserRepository struct {
	db *sql.DB
	c  *config.Config
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

	t := &models.Token{
		UserID: int64(u.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}

	jwtToken, err := t.Generate(ur.c.TokenKey)

	if err != nil {
		//TODO: Store to loger
		return err
	}

	q = `INSERT INTO users (registration_token) VALUES ($1) WHERE id = $2`

	if err = ur.db.QueryRow(q, jwtToken, u.ID).Err(); err != nil {
		return err
	}

	return nil

}
