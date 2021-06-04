package apiserver

func (s *server) updateBanksInstruments(bankiUrl string) {

	banks, err := s.instruments.Banks().GrabAll(bankiUrl)
	if err != nil {
		//TODO: Add to loger
		return
	}

	err = s.repository.BankInstrument().TruncateBanksInstruments()
	if err != nil {
		//TODO: Add to loger
		return
	}

	err = s.repository.BankInstrument().InsertBanksInstruments(banks)
	if err != nil {
		//TODO: Add to loger
		return
	}
}
