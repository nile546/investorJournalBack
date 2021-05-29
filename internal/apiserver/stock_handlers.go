package apiserver

func (s *server) renovStocks(spburl string, mskurl string) {

	stocks, err := s.instruments.Stocks().GrabAll(spburl, mskurl)
	if err != nil {
		//TODO: Add to loger
		return
	}
	s.repository.Stock().InsertStocks(stocks)
}
