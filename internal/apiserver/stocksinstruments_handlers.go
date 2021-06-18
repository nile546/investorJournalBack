package apiserver

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nile546/diplom/internal/models"
)

func (s *server) updateStocksInstruments(spburl string, mskurl string) {

	newStocks, err := s.instruments.Stocks().GrabAll(spburl, mskurl)
	if err != nil {
		s.logger.Errorf("Error grab stocks instruments: %+v", err)
		return
	}

	oldStocks, err := s.repository.StockInstrument().GetAllStockInstruments()
	if err != nil {
		s.logger.Errorf("Error fill stock_instruments: %+v", err)
		return
	}

	err = s.repository.StockInstrument().InsertStocksInstruments(getNewStockInstruments(*newStocks, *oldStocks))
	if err != nil {
		s.logger.Errorf("Error fill stock_instruments: %+v", err)
		return
	}
}

func getNewStockInstruments(newStocks []models.StockInstrument, oldStocks []models.StockInstrument) *[]models.StockInstrument {
	index := 0
	for _, oldStock := range oldStocks {
		for _, newStock := range newStocks {
			if *newStock.Ticker == *oldStock.Ticker {
				newStocks = append(newStocks[:index], newStocks[index+1:]...)
				oldStocks = append(oldStocks[:index], oldStocks[index+1:]...)
				break
			}
		}
	}
	return &newStocks
}

func (s *server) getPopularStockInstrument(w http.ResponseWriter, r *http.Request) {

	instrument, err := s.repository.StockInstrument().GetPopularStockInstrumentByUserID(s.session.userId)
	if err != nil {
		s.logger.Errorf("Error get popular stock instrument, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, instrument)

}

func (s *server) getPopularStockInstruments(w http.ResponseWriter, r *http.Request) {

	ids, err := s.repository.StockInstrument().GetPopularStockInstrumentsID()
	if ids == nil {
		s.error(w, "Stock deals not fount")
		return
	}
	if err != nil {
		s.logger.Errorf("Error get popular stock instruments, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil) //Добавить массив популярных акций

}

func (s *server) getStockInstrumentByID(w http.ResponseWriter, r *http.Request) {

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

	instrument, err := s.repository.StockInstrument().GetStockInstrumentByID(req.ID)
	if err != nil {
		s.logger.Errorf("Error get stock instrument by id, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, instrument)

}

func (s *server) getAllStockInstruments(w http.ResponseWriter, r *http.Request) {

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

	if err := s.repository.StockInstrument().GetAll(&req.TableParams); err != nil {
		s.logger.Errorf("Error get all stock instrument, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, req.TableParams)
}
