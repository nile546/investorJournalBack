package apiserver

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nile546/diplom/internal/models"
)

func (s *server) CreatePattern(w http.ResponseWriter, r *http.Request) {

	type request struct {
		Name           string                 `json:"name"`
		Description    string                 `json:"description"`
		UserID         int64                  `json:"user_id"`
		InstrumentType models.InstrumentTypes `json:"instrumentType"`
		Icon           string                 `json:"icon"`
	}

	req := &request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = validation.ValidateStruct(
		req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.UserID, validation.Required),
		validation.Field(&req.InstrumentType, validation.Required),
	)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = s.repository.Pattern().CreatePattern(&models.Pattern{
		Name:           req.Name,
		Description:    &req.Description,
		UserID:         &req.UserID,
		InstrumentType: req.InstrumentType,
		Icon:           &req.Icon,
	})
	if err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)
}

func (s *server) UpdatePattern(w http.ResponseWriter, r *http.Request) {

	type request struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	}

	req := &request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = validation.ValidateStruct(
		req,
		validation.Field(&req.ID, validation.Required),
		validation.Field(&req.Name, validation.Required),
	)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = s.repository.Pattern().UpdatePattern(&models.Pattern{
		ID:          req.ID,
		Name:        req.Name,
		Description: &req.Description,
		Icon:        &req.Icon,
	})
	if err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)
}

func (s *server) GetAllPattern(w http.ResponseWriter, r *http.Request) {
	var userID int64

	err := json.NewDecoder(r.Body).Decode(&userID)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	ptrns, err := s.repository.Pattern().GetAllPattern(userID)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, ptrns)
}

func (s *server) DeletePattern(w http.ResponseWriter, r *http.Request) {

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
