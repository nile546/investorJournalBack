package models

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int64  `json:"id"`
	Login             string `json:"login"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	EncryptedPassword string
	IsActive          bool      `json:"isActive"`
	CreatedAt         time.Time `json:"createdAt"`
	RegistrationToken string    `json:"registrationToken"`
}

func (u *User) Validate() error {

	if err := validation.ValidateStruct(
		u,
		validation.Field(&u.Login, validation.Required, validation.Length(3, 100)),
		validation.Field(&u.Email, validation.Required, is.Email, validation.Length(6, 100)),
		validation.Field(&u.Password, validation.Required, validation.Length(5, 100)),
	); err != nil {
		return err
	}

	return nil
}

func (u *User) EncryptPass() error {
	if u.Password == "" {
		//TODO: Move error message to constants
		return errors.New("Пароль не должен быть пустой строкой")
	}

	b, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.EncryptedPassword = string(b)
	u.Password = ""
	return nil
}
