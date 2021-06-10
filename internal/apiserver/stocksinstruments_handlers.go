package apiserver

func (s *server) updateStocksInstruments(spburl string, mskurl string) {

	stocks, err := s.instruments.Stocks().GrabAll(spburl, mskurl)
	if err != nil {
		//TODO: Add to loger
		return
	}

	err = s.repository.StockInstrument().TruncateStocksInstruments()
	if err != nil {
		//TODO: Add to loger
		return
	}

	err = s.repository.StockInstrument().InsertStocksInstruments(stocks)
	if err != nil {
		//TODO: Add to loger
		return
	}
}
