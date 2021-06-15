package apiserver

import "github.com/nile546/diplom/internal/models"

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
