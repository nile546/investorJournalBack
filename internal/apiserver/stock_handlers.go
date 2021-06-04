package apiserver

func (s *server) updateStocks(spburl string, mskurl string) {

	stocks, err := s.instruments.Stocks().GrabAll(spburl, mskurl)
	if err != nil {
		//TODO: Add to loger
		return
	}

	err = s.repository.Stock().TruncateStocks()
	if err != nil {
		//TODO: Add to loger
		return
	}

	err = s.repository.Stock().InsertStocks(stocks)
	if err != nil {
		//TODO: Add to loger
		return
	}
}
