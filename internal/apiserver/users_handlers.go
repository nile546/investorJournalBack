package apiserver

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
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
}
