package apiserver

func (s *server) renovCryptos(cryptoUrl string) {

	cryptos, err := s.instruments.Cryptos().GrabCrypto(cryptoUrl)
	if err != nil {
		//TODO: Add to loger
		return
	}
	s.repository.Crypto().InsertCrypto(cryptos)
}
