package apiserver

func (s *server) updateCryptoInstruments(cryptoUrl string, cryptoKey string) {

	cryptos, err := s.instruments.Cryptos().GrabAll(cryptoUrl, cryptoKey)
	if err != nil {
		s.logger.Errorf("Error grab crypto instruments: %+v", err)
		return
	}

	err = s.repository.CryptoInstrument().TruncateCryptoInstruments()
	if err != nil {
		s.logger.Errorf("Error truncate crypto_instruments: %+v", err)
		return
	}

	err = s.repository.CryptoInstrument().InsertCryptoInstruments(cryptos)
	if err != nil {
		s.logger.Errorf("Error fill crypto_instruments: %+v", err)
		return
	}
}
