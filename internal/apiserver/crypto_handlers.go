package apiserver

func (s *server) insertCryptos(cryptoUrl string, cryptoKey string) {

	cryptos, err := s.instruments.Cryptos().GrabAll(cryptoUrl, cryptoKey)
	if err != nil {
		//TODO: Add to loger
		return
	}

	err = s.repository.Crypto().InsertCrypto(cryptos)
	if err != nil {
		//TODO: Add to loger
		return
	}
}
