package apiserver

func (s *server) updateCryptoInstruments(cryptoUrl string, cryptoKey string) {

	cryptos, err := s.instruments.Cryptos().GrabAll(cryptoUrl, cryptoKey)
	if err != nil {
		//TODO: Add to loger
		return
	}

	err = s.repository.CryptoInstrument().TruncateCryptoInstruments()
	if err != nil {
		//TODO: Add to loger
		return
	}

	err = s.repository.CryptoInstrument().InsertCryptoInstruments(cryptos)
	if err != nil {
		//TODO: Add to loger
		return
	}
}
