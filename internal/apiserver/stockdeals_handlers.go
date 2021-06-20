package apiserver

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nile546/diplom/internal/models"
)

func (s *server) getAllStockDeals(w http.ResponseWriter, r *http.Request) {
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

	if err := s.repository.StockDeal().GetAll(&req.TableParams, s.session.userId); err != nil {
		s.logger.Errorf("Error get all stock deals, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, req.TableParams)

}

func (s *server) createStockDeal(w http.ResponseWriter, r *http.Request) {
	type request struct {
		StockDeal models.StockDeal `json:"stockDeal"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	if err := validation.ValidateStruct(
		req,
		validation.Field(&req.StockDeal, validation.Required),
	); err != nil {
		s.error(w, err.Error())
		return
	}

	err := s.repository.StockDeal().CreateStockDeal(&req.StockDeal, s.session.userId)
	if err != nil {
		s.logger.Errorf("Error create stock deal, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)
}

func (s *server) updateStockDeal(w http.ResponseWriter, r *http.Request) {

	type request struct {
		StockDeal models.StockDeal `json:"stockDeal"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	if err := validation.ValidateStruct(
		req,
		validation.Field(&req.StockDeal, validation.Required),
	); err != nil {
		s.error(w, err.Error())
		return
	}

	err := s.repository.StockDeal().UpdateStockDeal(&req.StockDeal, s.session.userId)
	if err != nil {
		s.logger.Errorf("Error update stock deal, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)

}

func (s *server) deleteStockDeal(w http.ResponseWriter, r *http.Request) {
	type request struct {
		ID int64 `json:"id"`
	}

	variability, err := s.repository.StockDeal().GetVariabilityByID(s.session.userId)
	if err != nil {
		s.logger.Errorf("Error get variability, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	if !variability {
		return
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

	err = s.repository.StockDeal().DeleteStockDeal(req.ID)
	if err != nil {
		s.logger.Errorf("Error delete stock deal, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)

}

func (s *server) getStockDealByID(w http.ResponseWriter, r *http.Request) {

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

	deal, err := s.repository.StockDeal().GetStockDealByID(req.ID)
	if err != nil {
		s.logger.Errorf("Error delete stock deal, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, deal)

}
