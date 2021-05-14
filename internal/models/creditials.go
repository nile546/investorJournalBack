package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Creditials struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	EncryptedPassword string
}

func (c *Creditials) EncryptPass() error {
	if c.Password == "" {
		//TODO: Move error message to constants
		return errors.New("Пароль не должен быть пустой строкой")
	}

	b, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	c.EncryptedPassword = string(b)
	return nil
}
