package apiserver

func (s *server) updateStocksInstruments(spburl string, mskurl string) {

	stocks, err := s.instruments.Stocks().GrabAll(spburl, mskurl)
	if err != nil {
		s.logger.Errorf("Error grab stocks instruments: %+v", err)
		return
	}

	//err = s.repository.StockInstrument().TruncateStocksInstruments()
	//if err != nil {
	//	s.logger.Errorf("Error truncate stock_instruments: %+v", err)
	//	return
	//}

	err = s.repository.StockInstrument().InsertStocksInstruments(stocks)
	if err != nil {
		s.logger.Errorf("Error fill stock_instruments: %+v", err)
		return
	}
}
