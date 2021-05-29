package apiserver

func (s *server) renovBanks(banksUrl string) {

	banks, err := s.instruments.Deposits().GrabBanks(banksUrl)
	if err != nil {
		//TODO: Add to loger
		return
	}
	s.repository.Bank().InsertBanks(banks)
}
