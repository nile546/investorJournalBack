package apiserver

func (s *server) updateCryptos(cryptoUrl string, cryptoKey string) {

	cryptos, err := s.instruments.Cryptos().GrabAll(cryptoUrl, cryptoKey)
	if err != nil {
		//TODO: Add to loger
		return
	}

	err = s.repository.Crypto().TruncateCrypto()
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
