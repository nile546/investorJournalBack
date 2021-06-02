package apiserver

func (s *server) insertBanks(bankiUrl string) {

	banks, err := s.instruments.Banks().GrabAll(bankiUrl)
	if err != nil {
		//TODO: Add to loger
		return
	}

	err = s.repository.Bank().InsertBanks(banks)
	if err != nil {
		//TODO: Add to loger
		return
	}
}
