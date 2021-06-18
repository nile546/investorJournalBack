package apiserver

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nile546/diplom/internal/models"
)

func (s *server) createPattern(w http.ResponseWriter, r *http.Request) {

	type request struct {
		Pattern *models.Pattern `json:"pattern"`
	}

	req := &request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = validation.ValidateStruct(
		req,
		validation.Field(&req.Pattern.Name, validation.Required),
		validation.Field(&req.Pattern.UserID, validation.Required),
		validation.Field(&req.Pattern.InstrumentType, validation.Required),
	)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = s.repository.Pattern().CreatePattern(req.Pattern)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)
}

func (s *server) updatePattern(w http.ResponseWriter, r *http.Request) {

	type request struct {
		Pattern *models.Pattern
	}

	req := &request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = validation.ValidateStruct(
		req,
		validation.Field(&req.Pattern.ID, validation.Required),
		validation.Field(&req.Pattern.Name, validation.Required),
		validation.Field(&req.Pattern.Description, validation.Required),
		validation.Field(&req.Pattern.Icon, validation.Required),
	)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = s.repository.Pattern().UpdatePattern(req.Pattern)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)
}

func (s *server) getAllPattern(w http.ResponseWriter, r *http.Request) {

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

	if err := s.repository.Pattern().GetAllPattern(&req.TableParams, s.session.userId); err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, req.TableParams)
}

func (s *server) deletePattern(w http.ResponseWriter, r *http.Request) {

	var id int64

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = s.repository.Pattern().DeletePattern(id)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)
}
