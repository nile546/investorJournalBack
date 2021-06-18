package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nile546/diplom/internal/models"
)

func (s *server) getAllDepositDeals(w http.ResponseWriter, r *http.Request) {
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

	if err := s.repository.DepositDeal().GetAll(&req.TableParams, s.session.userId); err != nil {
		s.logger.Errorf("Error get all deposit deal, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, req.TableParams)

}

func (s *server) createDepositDeal(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Deal models.DepositDeal `json:"depositDeal"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	if err := validation.ValidateStruct(
		req,
		validation.Field(&req.Deal.Bank, validation.Required),
		validation.Field(&req.Deal.EnterDateTime, validation.Required),
		validation.Field(&req.Deal.StartDeposit, validation.Required),
		validation.Field(&req.Deal.Percent, validation.Required),
		validation.Field(&req.Deal.UserID, validation.Required),
	); err != nil {
		s.error(w, err.Error())
		return
	}

	err := s.repository.DepositDeal().CreateDepositDeal(&req.Deal)
	if err != nil {
		s.logger.Errorf("Error create deposit deal, with error %+v", err)
		s.error(w, err.Error())
	}

	s.respond(w, nil)

}

func (s *server) updateDepositDeal(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Deal models.DepositDeal `json:"depositDeal"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	if req.Deal.UserID != s.session.userId {
		s.logger.Errorf("Error update deposit deal, with error %+v", errors.New("id user initiator does not match session user id"))
		s.error(w, "Id user initiator does not match session user id")
	}

	if err := validation.ValidateStruct(
		req,
		validation.Field(&req.Deal.ID, validation.Required),
		validation.Field(&req.Deal.EnterDateTime, validation.Required),
		validation.Field(&req.Deal.Bank, validation.Required),
		validation.Field(&req.Deal.StartDeposit, validation.Required),
		validation.Field(&req.Deal.Percent, validation.Required),
		validation.Field(&req.Deal.UserID, validation.Required),
	); err != nil {
		s.error(w, err.Error())
		return
	}

	err := s.repository.DepositDeal().UpdateDepositDeal(&req.Deal)
	if err != nil {
		s.logger.Errorf("Error update deposit deal, with error %+v", err)
		s.error(w, err.Error())
	}

	s.respond(w, nil)

}

func (s *server) deleteDepositDeal(w http.ResponseWriter, r *http.Request) {
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

	err := s.repository.DepositDeal().DeleteDepositDeal(req.ID)
	if err != nil {
		s.logger.Errorf("Error delete deposit deal, with error %+v", err)
		s.error(w, err.Error())
	}

	s.respond(w, nil)

}

func (s *server) getDepositDealByID(w http.ResponseWriter, r *http.Request) {

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

	deal, err := s.repository.DepositDeal().GetDepositDealByID(req.ID)
	if err != nil {
		s.logger.Errorf("Error delete deposit deal, with error %+v", err)
		s.error(w, err.Error())
	}

	s.respond(w, deal)

}
