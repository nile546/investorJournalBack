package apiserver

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nile546/diplom/internal/models"
)

func (s *server) CreateStockPattern(w http.ResponseWriter, r *http.Request) {

	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		UserID      int64  `json:"user_id"`
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
		validation.Field(&req.Description, validation.Required),
		validation.Field(&req.UserID, validation.Required),
	)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = s.repository.StockPattern().CreateStockPattern(&models.StockPattern{
		Name:        req.Name,
		Description: req.Description,
		UserID:      req.UserID,
	})
	if err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)
}

func (s *server) UpdateStockPattern(w http.ResponseWriter, r *http.Request) {

	type request struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
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
		validation.Field(&req.Description, validation.Required),
	)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = s.repository.StockPattern().UpdateStockPattern(&models.StockPattern{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
	}) //Обновлять время?
	if err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)
}

func (s *server) GetAllStockPattern(w http.ResponseWriter, r *http.Request) {

	var userID int64

	err := json.NewDecoder(r.Body).Decode(&userID)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	ptrns, err := s.repository.StockPattern().GetAllStockPattern(userID)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, ptrns)
}

func (s *server) DeleteStockPattern(w http.ResponseWriter, r *http.Request) {

	var id int64

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	err = s.repository.StockPattern().DeleteStockPattern(id)
	if err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil)
}
