package apiserver

func (s *server) updateBanksInstruments(bankiUrl string) {

	banks, err := s.instruments.Banks().GrabAll(bankiUrl)
	if err != nil {
		s.logger.Errorf("Error grab banks instruments: %+v", err)
		return
	}

	err = s.repository.BankInstrument().TruncateBanksInstruments()
	if err != nil {
		s.logger.Errorf("Error truncate banks_instruments: %+v", err)
		return
	}

	err = s.repository.BankInstrument().InsertBanksInstruments(banks)
	if err != nil {
		s.logger.Errorf("Error fill banks_instruments: %+v", err)
		return
	}
}
