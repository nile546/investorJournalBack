package apiserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/nile546/diplom/internal/models"
)

func (s *server) signup(w http.ResponseWriter, r *http.Request) {

	type request struct {
		Login    string `json:"login"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	if err := validation.ValidateStruct(
		req,
		validation.Field(&req.Login, validation.Required, validation.Length(3, 100)),
		validation.Field(&req.Email, validation.Required, is.Email, validation.Length(6, 100)),
		validation.Field(&req.Password, validation.Required, validation.Length(5, 100)),
	); err != nil {
		s.error(w, err.Error())
		return
	}

	u := &models.User{}
	u.Login = req.Login
	u.Email = req.Email
	u.Password = req.Password
	u.IsActive = false

	if err := s.repository.User().Create(u); err != nil {
		s.error(w, err.Error())
		return
	}

	t := &models.Token{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}

	jwtToken, err := t.Generate(tokenKey)
	u.RegistrationToken = jwtToken

	err = s.repository.User().Update(u)

	s.error(w, err.Error())

}
