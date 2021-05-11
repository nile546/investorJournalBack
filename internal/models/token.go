package models

import "github.com/dgrijalva/jwt-go"

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
