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

func (s *server) updateSession(w http.ResponseWriter, r *http.Request) {

	type request struct {
		RT string `json:"refreshToken"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	if err := validation.ValidateStruct(req,
		validation.Field(&req.RT, validation.Required, is.UUIDv4),
	); err != nil {
		s.error(w, err.Error())
		return
	}

	refreshToken, userID, err := s.repository.User().UpdateRefreshToken(req.RT)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	at := &models.Token{
		UserID: userID,
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

	rtc := &http.Cookie{
		Name:     "rt",
		Value:    refreshToken,
		HttpOnly: true,
		Path:     apiRoute + updateSessionRoute,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	if production {
		rtc.Secure = true
		rtc.SameSite = http.SameSiteStrictMode
	}

	http.SetCookie(w, rtc)

	s.respond(w, nil)
}

func (s *server) clearSession(w http.ResponseWriter, r *http.Request) {

	atc, err := r.Cookie("at")

	if err != nil {
		s.respond(w, err.Error())
		return
	}

	clearAtc := &http.Cookie{
		Name:     "at",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	}
	http.SetCookie(w, clearAtc)

	clearRtc := &http.Cookie{
		Name:     "rt",
		Value:    "",
		Path:     apiRoute + updateSessionRoute,
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	}
	http.SetCookie(w, clearRtc)

	accessToken := &models.Token{}
	if err := accessToken.GetClaims(atc.Value, tokenKey); err != nil {
		s.respond(w, err.Error())
		return
	}

	u := &models.User{
		ID: accessToken.UserID,
	}

	if err = s.repository.User().DeleteRefreshTokenByUser(u); err != nil {
		s.logger.Errorf("Error delete from refresh_token by id: %+v", err)
	}

	s.respond(w, nil)

}
