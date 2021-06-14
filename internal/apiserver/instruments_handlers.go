package apiserver

import (
	"time"
)

type updateInstrumentsConfig struct {
	spbExchangeUrl string
	mskStocksUrl   string
	bankiUrl       string
	cryptoUrl      string
	cryptoKey      string
	hours          int
	minutes        int
	seconds        int
}

func (s *server) updateInstruments(callHandlers func(c *updateInstrumentsConfig), c *updateInstrumentsConfig) error {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}

	now := time.Now().Local()
	firstCallTime := time.Date(
		now.Year(), now.Month(), now.Day(), c.hours, c.minutes, c.seconds, 0, loc)
	if firstCallTime.Before(now) {
		firstCallTime = firstCallTime.Add(time.Hour * 24)
	}

	duration := firstCallTime.Sub(time.Now().Local())

	go func() {
		time.Sleep(duration)
		for {
			callHandlers(c)
			time.Sleep(time.Hour * 24)
		}
	}()

	return nil
}

func (s *server) callUpdateHandlers(c *updateInstrumentsConfig) {
	s.updateStocksInstruments(c.spbExchangeUrl, c.mskStocksUrl)
	s.updateCryptoInstruments(c.cryptoUrl, c.cryptoKey)
	s.updateBanksInstruments(c.bankiUrl)
}
