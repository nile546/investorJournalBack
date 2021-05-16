package apiserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
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

	u := &models.User{
		Login:    req.Login,
		Email:    req.Email,
		Password: req.Password,
		IsActive: false,
	}

	if err := u.EncryptPass(); err != nil {
		s.error(w, err.Error())
	}

	if err := s.repository.User().Create(u); err != nil {
		s.error(w, err.Error())
		return
	}

	t := template.New("signup-email.html")

	t, err := t.ParseFiles("tmpl/signup-email.html")
	if err != nil {
		s.error(w, err.Error())
		return
	}

	rt := &models.Token{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}

	jwtToken, err := rt.Generate(tokenKey)
	if err != nil {
		return
	}

	u.RegistrationToken = jwtToken

	err = s.repository.User().Update(u)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	buf := new(bytes.Buffer)
	data := struct {
		Login    string
		RegToken string
		URLPath  string
		Address  string
	}{
		Login:    u.Login,
		RegToken: u.RegistrationToken,
		URLPath:  confirmSignupRoute,
		Address:  addr,
	}

	if err = t.Execute(buf, data); err != nil {
		s.error(w, err.Error())
		return
	}

	if err := s.mailer.Send([]string{u.Email}, buf.String()); err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)

}

func (s *server) signin(w http.ResponseWriter, r *http.Request) {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := &request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = validation.ValidateStruct(
		req,
		validation.Field(&req.Email, validation.Required, is.Email, validation.Length(6, 100)),
		validation.Field(&req.Password, validation.Required, validation.Length(5, 100)),
	)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	c := &models.Creditials{
		Email:    req.Email,
		Password: req.Password,
	}

	err = c.EncryptPass()
	if err != nil {
		s.error(w, err.Error())
		return
	}

	u, err := s.repository.User().GetUserByEmail(c.Email)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	if u.EncryptedPassword != c.EncryptedPassword {
		err = errors.New("Invalid password!")
		s.error(w, err.Error())
		return
	}

	if u.IsActive == false {
		err = errors.New("Account not verified!")
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)

}
