package apiserver

import (
	"encoding/json"
	"errors"
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

	if err := s.repository.CryptoDeal().GetAll(&req.TableParams); err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, req.TableParams)

}

func (s *server) createCryptoDeal(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Deal models.CryptoDeal `json:"cryptoDeal"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	if err := validation.ValidateStruct(
		req,
		validation.Field(&req.Deal, validation.Required),
	); err != nil {
		s.error(w, err.Error())
		return
	}

	err := s.repository.CryptoDeal().CreateCryptoDeal(&req.Deal)
	if err != nil {
		s.logger.Errorf("Error create сrypto deal, with error %+v", err)
		s.error(w, err.Error())
	}

}

func (s *server) updateCryptoDeal(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Deal models.CryptoDeal `json:"cryptoDeal"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	if req.Deal.UserID != s.session.userId {
		s.logger.Errorf("Error update crypto deal, with error %+v", errors.New("id user initiator does not match session user id"))
		s.error(w, "Id user initiator does not match session user id")
	}

	if err := validation.ValidateStruct(
		req,
		validation.Field(&req.Deal, validation.Required),
	); err != nil {
		s.error(w, err.Error())
		return
	}

	err := s.repository.CryptoDeal().UpdateCryptoDeal(&req.Deal)
	if err != nil {
		s.logger.Errorf("Error update crypto deal, with error %+v", err)
		s.error(w, err.Error())
	}

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
