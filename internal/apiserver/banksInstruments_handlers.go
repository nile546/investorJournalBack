package apiserver

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nile546/diplom/internal/models"
)

func (s *server) updateBanksInstruments(bankiUrl string) {

	newBanks, err := s.instruments.Banks().GrabAll(bankiUrl)
	if err != nil {
		s.logger.Errorf("Error grab banks instruments: %+v", err)
		return
	}

	oldBanks, err := s.repository.BankInstrument().GetAllBankInstruments()
	if err != nil {
		s.logger.Errorf("Error fill banks_instruments: %+v", err)
		return
	}

	err = s.repository.BankInstrument().InsertBanksInstruments(getNewBankInstruments(*newBanks, *oldBanks))
	if err != nil {
		s.logger.Errorf("Error fill banks_instruments: %+v", err)
		return
	}
}

func getNewBankInstruments(newBanks []models.BankInstrument, oldBanks []models.BankInstrument) *[]models.BankInstrument {
	index := 0
	for _, oldBank := range oldBanks {
		for _, newBank := range newBanks {
			if newBank.Title == oldBank.Title {
				newBanks = append(newBanks[:index], newBanks[index+1:]...)
				oldBanks = append(oldBanks[:index], oldBanks[index+1:]...)
				break
			}
		}
	}
	return &newBanks
}

func (s *server) getPopularBankInstrument(w http.ResponseWriter, r *http.Request) {

	instrument, err := s.repository.BankInstrument().GetPopularBankInstrumentByUserID(s.session.userId)
	if err != nil {
		s.logger.Errorf("Error get popular bank instrument, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, instrument)

}

func (s *server) getPopularBankInstruments(w http.ResponseWriter, r *http.Request) {

	ids, err := s.repository.BankInstrument().GetPopularBankInstrumentsID()
	if ids == nil {
		s.error(w, "Bank deals not fount")
		return
	}
	if err != nil {
		s.logger.Errorf("Error get popular bank instruments, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil) //Добавить массив популярных банков

}

func (s *server) getBankInstrumentByID(w http.ResponseWriter, r *http.Request) {

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

	instrument, err := s.repository.BankInstrument().GetBankInstrumentByID(req.ID)
	if err != nil {
		s.logger.Errorf("Error get bank instrument by id, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, instrument)

}

func (s *server) getAllBankInstruments(w http.ResponseWriter, r *http.Request) {

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

	if err := s.repository.BankInstrument().GetAll(&req.TableParams); err != nil {
		s.logger.Errorf("Error get all bank instrument, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, req.TableParams)
}
