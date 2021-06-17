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

	if err := s.repository.StockDeal().GetAll(&req.TableParams); err != nil {
		s.error(w, err.Error())
		return
	}

	s.respond(w, req.TableParams)

}

func (s *server) CreateStockDeal(w http.ResponseWriter, r *http.Request) {

	type request struct {
		Deal models.StockDeal `json:"stock_deal"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, err.Error())
		return
	}

	if err := validation.ValidateStruct(
		req,
		validation.Field(&req.Deal.EnterDateTime, validation.Required),
		//TODO: Need to check required pageNumber and itemsPerPage
	); err != nil {
		s.error(w, err.Error())
		return
	}

	//err := s.repository.StockDeal().CreateStockDeal(&req.Deal)
	//if err != nil{
	//	s.logger.Errorf("Error create stock deal, with error")
	//	s.error(w, err)
	//}

}
