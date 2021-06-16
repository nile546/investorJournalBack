package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int64      `json:"id"`
	Login             string     `json:"login"`
	Email             string     `json:"email"`
	Password          string     `json:"password"`
	EncryptedPassword string     `json:"-"`
	IsActive          bool       `json:"isActive"`
	CreatedAt         time.Time  `json:"createdAt"`
	DateGrab          *time.Time `json:"dateGrab"`
	AutoGrabDeals     bool       `json:"autoGrabDeals"`
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

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))

	return err == nil
}
