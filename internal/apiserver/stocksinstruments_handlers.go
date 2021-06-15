package apiserver

import (
	"github.com/nile546/diplom/internal/models"
)

func (s *server) updateStocksInstruments(spburl string, mskurl string) {

	newStocks, err := s.instruments.Stocks().GrabAll(spburl, mskurl)
	if err != nil {
		s.logger.Errorf("Error grab stocks instruments: %+v", err)
		return
	}

	oldStocks, err := s.repository.StockInstrument().GetAllStockInstruments()
	if err != nil {
		s.logger.Errorf("Error fill stock_instruments: %+v", err)
		return
	}

	err = s.repository.StockInstrument().InsertStocksInstruments(getNewStockInstruments(*newStocks, *oldStocks))
	if err != nil {
		s.logger.Errorf("Error fill stock_instruments: %+v", err)
		return
	}
}

func getNewStockInstruments(newStocks []models.StockInstrument, oldStocks []models.StockInstrument) *[]models.StockInstrument {
	index := 0
	for _, oldStock := range oldStocks {
		for _, newStock := range newStocks {
			if *newStock.Ticker == *oldStock.Ticker {
				newStocks = append(newStocks[:index], newStocks[index+1:]...)
				oldStocks = append(oldStocks[:index], oldStocks[index+1:]...)
				break
			}
		}
	}
	return &newStocks
}
