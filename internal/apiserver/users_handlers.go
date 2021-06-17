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

	if err := s.repository.User().CreateUser(u); err != nil {
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

	buf := new(bytes.Buffer)
	data := struct {
		Login    string
		RegToken string
		URLPath  string
		Address  string
	}{
		Login:    u.Login,
		RegToken: jwtToken,
		URLPath:  landingRoute + confirmSignupRoute,
		Address:  protocol + addrLand,
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

func (s *server) confirmSignup(w http.ResponseWriter, r *http.Request) {

	type request struct {
		Token string `json:"token"`
	}

	req := &request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = validation.ValidateStruct(
		req,
		validation.Field(&req.Token, validation.Required),
	)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	tkn := &models.Token{}
	err = tkn.GetClaims(req.Token, tokenKey)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	if tkn.StandardClaims.ExpiresAt < time.Now().Unix() {
		err = errors.New("Время действия токена истекло")
		s.error(w, err.Error())
		return
	}

	err = s.repository.User().UpdateIsActiveByUserID(tkn.UserID)
	if err != nil {
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

	c := &models.Credentials{
		Email:    req.Email,
		Password: req.Password,
	}

	u, err := s.repository.User().GetUserByEmail(c.Email)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	if !u.ComparePassword(c.Password) {
		err = errors.New("Invalid password!")
		s.error(w, err.Error())
		return
	}

	if u.IsActive == false {
		err = errors.New("Account not verified!")
		s.error(w, err.Error())
		return
	}

	at := &models.Token{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	accessToken, err := at.Generate(tokenKey)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	atc := &http.Cookie{
		Name:     "at",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Minute * 35),
	}

	if production {
		atc.Secure = true
		atc.SameSite = http.SameSiteStrictMode
	}

	http.SetCookie(w, atc)

	refreshToken, err := s.repository.User().SetRefreshToken(u.ID)

	rtc := &http.Cookie{
		Name:     "rt",
		Value:    refreshToken,
		HttpOnly: true,
		Path:     apiRoute + authRoute + refreshRoute,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	if production {
		rtc.Secure = true
		rtc.SameSite = http.SameSiteStrictMode
	}

	http.SetCookie(w, rtc)

	if u.AutoGrabDeals {
		token, err := s.repository.TinkoffToken().GetTinkoffToken(u.ID)
		if err != nil {
			s.logger.Errorf("Error insert Tinkoff stock deals, with error: %+v", err)
		}

		s.getTinkoffStockDeals(token, u.AutoGrabDeals)

	}

	s.respond(w, u)

}

func (s *server) getUser(w http.ResponseWriter, r *http.Request) {

	type request struct {
		ID int64 `json:"id"`
	}

	req := &request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		u, err := s.repository.User().GetUserByID(s.session.userId)
		if err != nil {
			s.error(w, err.Error())
			return
		}

		s.respond(w, u)
		return
	}
	u, err := s.repository.User().GetUserByID(req.ID)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, u)

}
