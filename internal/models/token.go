package models

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserID int64 `json:"userId"`
	jwt.StandardClaims
}

func (t *Token) Generate(tokenKey string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	signesToken, err := token.SignedString([]byte(tokenKey))

	if err != nil {
		return "", nil
	}

	return signesToken, nil
}

func (t *Token) GetClaims(tokenString string, tokenKey string) error {
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("Проверка токена на подлинность не пройдена")
	}

	claims, ok := token.Claims.(*Token)
	if !ok {
		return errors.New("Время жизни токена истекло")
	}

	t.UserID = claims.UserID
	t.StandardClaims.ExpiresAt = claims.StandardClaims.ExpiresAt

	return nil
}
