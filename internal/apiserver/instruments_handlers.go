package apiserver

import (
	"time"
)

type instrumentsConfig struct {
	spbExchangeUrl string
	mskStocksUrl   string
	bankiUrl       string
	cryptoUrl      string
	cryptoKey      string
}

func (s *server) updateInstruments(hour, min, sec int, callHandlers func(c *instrumentsConfig), iC *instrumentsConfig) error {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}

	now := time.Now().Local()
	firstCallTime := time.Date(
		now.Year(), now.Month(), now.Day(), hour, min, sec, 0, loc)
	if firstCallTime.Before(now) {
		firstCallTime = firstCallTime.Add(time.Hour * 24)
	}

	duration := firstCallTime.Sub(time.Now().Local())

	go func() {
		time.Sleep(duration)
		for {
			callHandlers(iC)
			time.Sleep(time.Hour * 24)
		}
	}()

	return nil
}

func (s *server) callUpdateHandlers(c *instrumentsConfig) {
	s.updateStocks(c.spbExchangeUrl, c.mskStocksUrl)
	s.updateCryptos(c.cryptoUrl, c.cryptoKey)
	s.updateBanks(c.bankiUrl)
}
