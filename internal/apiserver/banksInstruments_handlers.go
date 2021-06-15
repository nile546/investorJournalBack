package apiserver

import "github.com/nile546/diplom/internal/models"

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
