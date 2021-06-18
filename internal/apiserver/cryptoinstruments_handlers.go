package apiserver

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nile546/diplom/internal/models"
)

func (s *server) updateCryptoInstruments(cryptoUrl string, cryptoKey string) {

	newCryptos, err := s.instruments.Cryptos().GrabAll(cryptoUrl, cryptoKey)
	if err != nil {
		s.logger.Errorf("Error grab crypto instruments: %+v", err)
		return
	}

	oldCryptos, err := s.repository.CryptoInstrument().GetAllCryptoInstruments()
	if err != nil {
		s.logger.Errorf("Error fill crypto_instruments: %+v", err)
		return
	}

	err = s.repository.CryptoInstrument().InsertCryptoInstruments(getNewCryptoInstruments(*newCryptos, *oldCryptos))
	if err != nil {
		s.logger.Errorf("Error fill crypto_instruments: %+v", err)
		return
	}
}

func getNewCryptoInstruments(newCrypto []models.CryptoInstrument, oldCrypto []models.CryptoInstrument) *[]models.CryptoInstrument {
	index := 0
	for _, oldCrypt := range oldCrypto {
		for _, newCrypt := range newCrypto {
			if newCrypt.Title == oldCrypt.Title {
				newCrypto = append(newCrypto[:index], newCrypto[index+1:]...)
				oldCrypto = append(oldCrypto[:index], oldCrypto[index+1:]...)
				break
			}
		}
	}
	return &newCrypto
}

func (s *server) getPopularCryptoInstrument(w http.ResponseWriter, r *http.Request) {

	instrument, err := s.repository.CryptoInstrument().GetPopularCryptoInstrumentByUserID(s.session.userId)
	if err != nil {
		s.logger.Errorf("Error get popular crypto instrument, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, instrument)

}

func (s *server) getPopularCryptoInstruments(w http.ResponseWriter, r *http.Request) {

	ids, err := s.repository.CryptoInstrument().GetPopularCryptoInstrumentsID()
	if ids == nil {
		s.error(w, "Crypto deals not fount")
		return
	}
	if err != nil {
		s.logger.Errorf("Error get popular crypto instruments, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, nil) //Добавить массив популярных криптовалют
}

func (s *server) getCryptoInstrumentByID(w http.ResponseWriter, r *http.Request) {

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

	instrument, err := s.repository.CryptoInstrument().GetCryptoInstrumentByID(req.ID)
	if err != nil {
		s.logger.Errorf("Error get crypto instrument by id, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, instrument)

}

func (s *server) getAllCryptoInstruments(w http.ResponseWriter, r *http.Request) {

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

	if err := s.repository.CryptoInstrument().GetAll(&req.TableParams); err != nil {
		s.logger.Errorf("Error get all crypto instruments, with error %+v", err)
		s.error(w, err.Error())
		return
	}

	s.respond(w, req.TableParams)
}
