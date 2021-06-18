package apiserver

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nile546/diplom/internal/models"
)

func (s *server) createStrategy(w http.ResponseWriter, r *http.Request) {

	type request struct {
		Strategy *models.Strategy `json:"strategy"`
	}

	req := &request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = validation.ValidateStruct(
		req,
		validation.Field(&req.Strategy.Name, validation.Required),
		validation.Field(&req.Strategy.UserID, validation.Required),
		validation.Field(&req.Strategy.InstrumentType, validation.Required),
	)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = s.repository.Strategy().CreateStrategy(req.Strategy)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)
}

func (s *server) updateStockStrategy(w http.ResponseWriter, r *http.Request) {

	type request struct {
		Strategy *models.Strategy `json:"strategy"`
	}

	req := &request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = validation.ValidateStruct(
		req,
		validation.Field(&req.Strategy.ID, validation.Required),
		validation.Field(&req.Strategy.Name, validation.Required),
	)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = s.repository.Strategy().UpdateStrategy(req.Strategy)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)
}

func (s *server) getAllStrategy(w http.ResponseWriter, r *http.Request) {

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

	if err := s.repository.Strategy().GetAllStrategy(&req.TableParams, s.session.userId); err != nil {
		s.logger.Errorf("Error get all strategy, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, req.TableParams)

}

func (s *server) deleteStrategy(w http.ResponseWriter, r *http.Request) {

	type request struct {
		ID int64 `json:"id"`
	}

	req := &request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
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

	err = s.repository.Strategy().DeleteStrategy(req.ID)
	if err != nil {
		s.error(w, err.Error())
		s.logger.Errorf("Error delete strategy, with error %+v", err)
		return
	}

	s.respond(w, nil)
}
