package apiserver

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nile546/diplom/internal/models"
)

func (s *server) getAllCryptoDeals(w http.ResponseWriter, r *http.Request) {
	type request struct {
		TableParams models.TableParams `json:"tableParams"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	if err := validation.ValidateStruct(
		req,
		validation.Field(&req.TableParams, validation.Required),
	); err != nil {
		s.error(w, err.Error())
		return
	}

	if err := s.repository.CryptoDeal().GetAll(&req.TableParams, s.session.userId); err != nil {
		s.logger.Errorf("Error get all crypto deals, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, req.TableParams)

}

func (s *server) createCryptoDeal(w http.ResponseWriter, r *http.Request) {
	type request struct {
		CryptoDeal models.CryptoDeal `json:"cryptoDeal"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	if err := validation.ValidateStruct(
		req,
		validation.Field(&req.CryptoDeal, validation.Required),
	); err != nil {
		s.error(w, err.Error())
		return
	}

	err := s.repository.CryptoDeal().CreateCryptoDeal(&req.CryptoDeal, s.session.userId)
	if err != nil {
		s.logger.Errorf("Error create сrypto deal, with error %+v", err)
		s.error(w, err.Error())
	}

	s.respond(w, nil)

}

func (s *server) updateCryptoDeal(w http.ResponseWriter, r *http.Request) {
	type request struct {
		CryptoDeal models.CryptoDeal `json:"cryptoDeal"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	if err := validation.ValidateStruct(
		req,
		validation.Field(&req.CryptoDeal, validation.Required),
	); err != nil {
		s.error(w, err.Error())
		return
	}

	err := s.repository.CryptoDeal().UpdateCryptoDeal(&req.CryptoDeal, s.session.userId)
	if err != nil {
		s.logger.Errorf("Error update crypto deal, with error %+v", err)
		s.error(w, err.Error())
	}

	s.respond(w, nil)

}

func (s *server) deleteCryptoDeal(w http.ResponseWriter, r *http.Request) {
	type request struct {
		ID int64 `json:"id"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	if err := validation.ValidateStruct(
		req,
		validation.Field(&req.ID, validation.Required),
	); err != nil {
		s.error(w, err.Error())
		return
	}

	err := s.repository.CryptoDeal().DeleteCryptoDeal(req.ID)
	if err != nil {
		s.logger.Errorf("Error delete crypto deal, with error %+v", err)
		s.error(w, err.Error())
	}

	s.respond(w, nil)

}

func (s *server) getCryptoDealByID(w http.ResponseWriter, r *http.Request) {

	type request struct {
		ID int64 `json:"id"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	if err := validation.ValidateStruct(
		req,
		validation.Field(&req.ID, validation.Required),
	); err != nil {
		s.error(w, err.Error())
		return
	}

	deal, err := s.repository.CryptoDeal().GetCryptoDealByID(req.ID)
	if err != nil {
		s.logger.Errorf("Error delete crypto deal, with error %+v", err)
		s.error(w, err.Error())
	}

	s.respond(w, deal)

}
